package my_coredns_plugin

import (
	"context"
	"fmt"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

type MyPlugin struct {
	Next plugin.Handler
}

func (m MyPlugin) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	fmt.Println("=====Yoooo Serving DNS")
	pw := NewResponsePrinter(w)
	return plugin.NextOrFailure(m.Name(), m.Next, ctx, pw, r)
}

func (m MyPlugin) Name() string {
	return "myplugin"
}

type ResponsePrinter struct {
	dns.ResponseWriter
}

// NewResponsePrinter returns ResponseWriter.
func NewResponsePrinter(w dns.ResponseWriter) *ResponsePrinter {
	return &ResponsePrinter{ResponseWriter: w}
}
