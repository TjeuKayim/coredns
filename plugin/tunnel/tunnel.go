// Package tunnel implements a plugin that returns details about the resolving
// querying it.
package tunnel

import (
	"context"

	"github.com/coredns/coredns/request"

	"github.com/miekg/dns"
)

const name = "tunnel"

// Tunnel is a plugin that returns your IP address, port and the protocol used for connecting
// to CoreDNS.
type Tunnel struct{}

// ServeDNS implements the plugin.Handler interface.
func (wh Tunnel) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	a := new(dns.Msg)
	a.SetReply(r)
	a.Authoritative = true

	rr := new(dns.CNAME)
	rr.Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeCNAME, Class: state.QClass()}

	name := state.QName()
	idx := dns.Split(name)
	firstLabel := name[:idx[1]]
	zone := name[idx[1]:]
	rr.Target = reverse(firstLabel) + zone

	a.Answer = []dns.RR{rr}

	w.WriteMsg(a)

	return 0, nil
}

// Name implements the Handler interface.
func (wh Tunnel) Name() string { return name }

func reverse(name string) (reversed string) {
	for _, v := range name {
		reversed = string(v) + reversed
	}
	return
}
