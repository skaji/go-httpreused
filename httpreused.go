package httpreused

import (
	"net"
	"net/http"
	"net/http/httptrace"
)

type roundTripper struct {
	base http.RoundTripper
}

func (rt *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	reused := "false"
	remoteIP := ""
	ctx := httptrace.WithClientTrace(req.Context(), &httptrace.ClientTrace{
		GotConn: func(c httptrace.GotConnInfo) {
			if c.Reused {
				reused = "true"
			}
			remoteIP, _, _ = net.SplitHostPort(c.Conn.RemoteAddr().String())
		},
	})
	res, err := rt.base.RoundTrip(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	res.Header.Set("X-Connection-Reused", reused)
	res.Header.Set("X-Connection-IP", remoteIP)
	return res, err
}

func Wrap(c *http.Client) *http.Client {
	if c.Transport == nil {
		c.Transport = http.DefaultTransport
	}
	b := c.Transport
	if wrapped, ok := b.(*roundTripper); ok {
		b = wrapped.base
	}
	c.Transport = &roundTripper{
		base: b,
	}
	return c
}
