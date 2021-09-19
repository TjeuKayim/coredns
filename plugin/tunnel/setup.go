package tunnel

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("tunnel", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'tunnel'
	if c.NextArg() {
		return plugin.Error("tunnel", c.ArgErr())
	}

	t := Tunnel{}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		t.Next = next
		return t
	})

	return nil
}
