// Package tunnel implements a plugin that returns details about the resolving
// querying it.
package tunnel

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"

	"github.com/miekg/dns"
)

const name = "tunnel"

// Tunnel is a plugin that returns your IP address, port and the protocol used for connecting
// to CoreDNS.
type Tunnel struct {
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface.
func (wh Tunnel) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	if r.Question[0].Qtype != dns.TypeCNAME {
		return plugin.NextOrFailure(wh.Name(), wh.Next, ctx, w, r)
	}

	state := request.Request{W: w, Req: r}

	a := new(dns.Msg)
	a.SetReply(r)
	a.Authoritative = true

	rr := new(dns.CNAME)
	rr.Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeCNAME, Class: state.QClass()}

	name := state.QName()
	idx := dns.Split(name)
	i := idx[len(idx)-2] - 1
	firstLabel := name[:i]
	zone := name[i:]
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
