package ssh

import (
	"bytes"
	"context"
	"encoding/base64"
	core "github.com/v2fly/v2ray-core/v4"
	"github.com/v2fly/v2ray-core/v4/common"
	"github.com/v2fly/v2ray-core/v4/common/buf"
	"github.com/v2fly/v2ray-core/v4/common/net"
	"github.com/v2fly/v2ray-core/v4/common/retry"
	"github.com/v2fly/v2ray-core/v4/common/session"
	"github.com/v2fly/v2ray-core/v4/common/signal"
	"github.com/v2fly/v2ray-core/v4/common/task"
	"github.com/v2fly/v2ray-core/v4/features/policy"
	"github.com/v2fly/v2ray-core/v4/proxy"
	"github.com/v2fly/v2ray-core/v4/transport"
	"github.com/v2fly/v2ray-core/v4/transport/internet"
	"golang.org/x/crypto/ssh"
	"strings"
	"sync"
)

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		c := &Client{}
		return c, core.RequireFeatures(ctx, func(policyManager policy.Manager) error {
			return c.Init(config.(*Config), policyManager)
		})
	}))
}

var _ proxy.Outbound = (*Client)(nil)
var _ common.Closable = (*Client)(nil)

type Client struct {
	sync.Mutex
	sessionPolicy   policy.Session
	server          net.Destination
	client          *ssh.Client
	username        string
	auth            []ssh.AuthMethod
	hostKeyCallback ssh.HostKeyCallback
}

func (c *Client) Init(config *Config, policyManager policy.Manager) error {
	c.sessionPolicy = policyManager.ForLevel(config.UserLevel)
	c.server = net.Destination{
		Network: net.Network_TCP,
		Address: config.Address.AsAddress(),
		Port:    net.Port(config.Port),
	}
	c.username = config.User
	if c.username == "" {
		c.username = "root"
	}

	if config.PrivateKey != "" {
		var signer ssh.Signer
		var err error
		if config.Password == "" {
			signer, err = ssh.ParsePrivateKey([]byte(config.PrivateKey))
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(config.PrivateKey), []byte(config.Password))
		}
		if err != nil {
			return newError("parse private key").Base(err)
		}
		c.auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	} else if config.Password != "" {
		c.auth = []ssh.AuthMethod{ssh.Password(config.Password)}
	}

	var keys []ssh.PublicKey
	if config.PublicKey != "" {
		for _, str := range strings.Split(config.PublicKey, "\n") {
			str = strings.TrimSpace(str)
			if str == "" {
				continue
			}
			key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(str))
			if err != nil {
				if err != nil {
					return newError(err, "parse public key").Base(err)
				}
			}
			keys = append(keys, key)
		}
	}
	if keys != nil {
		c.hostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			for _, pk := range keys {
				if bytes.Equal(key.Marshal(), pk.Marshal()) {
					return nil
				}
			}
			return newError("ssh: host key mismatch, server send ", key.Type(), " ", base64.StdEncoding.EncodeToString(key.Marshal()))
		}
	} else {
		c.hostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			newError("ssh: server send ", key.Type(), " ", base64.StdEncoding.EncodeToString(key.Marshal())).AtInfo().WriteToLog()
			return nil
		}
	}
	return nil
}

func (c *Client) Process(ctx context.Context, link *transport.Link, dialer internet.Dialer) error {
	outbound := session.OutboundFromContext(ctx)
	if outbound == nil || !outbound.Target.IsValid() {
		return newError("target not specified")
	}
	destination := outbound.Target
	network := destination.Network
	if network != net.Network_TCP {
		return newError("only TCP is supported in SSH proxy")
	}

	sc := c.client
	if sc == nil {
		c.Lock()
		sc = c.client
		if c.client == nil {
			client, err := c.connect(ctx, dialer)
			if err != nil {
				return err
			}
			go func() {
				err = client.Wait()
				if err != nil {
					newError("ssh client closed").Base(err).AtInfo().WriteToLog()
				}
				c.Lock()
				c.client = nil
				c.Unlock()
			}()
			sc = client
		}
		c.Unlock()
	}

	conn, err := sc.Dial("tcp", destination.NetAddr())
	if err != nil {
		return newError("failed to open ssh proxy connection").Base(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	timer := signal.CancelAfterInactivity(ctx, cancel, c.sessionPolicy.Timeouts.ConnectionIdle)

	if err := task.Run(ctx, func() error {
		defer timer.SetTimeout(c.sessionPolicy.Timeouts.DownlinkOnly)
		return buf.Copy(link.Reader, buf.NewWriter(conn), buf.UpdateActivity(timer))
	}, func() error {
		defer timer.SetTimeout(c.sessionPolicy.Timeouts.UplinkOnly)
		return buf.Copy(buf.NewReader(conn), link.Writer, buf.UpdateActivity(timer))
	}); err != nil {
		return newError("connection ends").Base(err)
	}

	return nil

}

func (c *Client) connect(ctx context.Context, dialer internet.Dialer) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User:            c.username,
		Auth:            c.auth,
		HostKeyCallback: c.hostKeyCallback,
	}

	var conn internet.Connection
	err := retry.ExponentialBackoff(2, 100).On(func() error {
		rawConn, err := dialer.Dial(ctx, c.server)
		if err != nil {
			return err
		}
		conn = rawConn
		return nil
	})

	if err != nil {
		return nil, newError("failed to connect to destination").AtWarning().Base(err)
	}
	clientConn, chans, reqs, err := ssh.NewClientConn(conn, c.server.Address.String(), config)
	if err != nil {
		return nil, newError("failed to ssh").Base(err)
	}
	client := ssh.NewClient(clientConn, chans, reqs)
	c.client = client
	return client, nil
}

func (c *Client) Close() error {
	sc := c.client
	if sc != nil {
		return sc.Close()
	}
	return nil
}
