package proxy

import (
	"net/http"
	"net/url"
	"sync/atomic"
)

type FuncProxy func(r *http.Request) (*url.URL, error)

type roundRobinSwitcher struct {
	proxyURLs []*url.URL
	index     uint32
}

func (r *roundRobinSwitcher) GetProxy(pr *http.Request) (*url.URL, error) {
	index := atomic.AddUint32(&r.index, 1) - 1
	u := r.proxyURLs[index%uint32(len(r.proxyURLs))]
	return u, nil
}

// RoundRobinProxySwitcher create a proxy switch function which rotates ProxyURLs on every request.
// The proxy type is determined by the URL scheme. "http", "https" and "socks5" are supported.
// If the scheme is empty, "http" is assumed.
/*func RoundRobinProxySwitcher(ProxyURLs ...string) (FuncProxy, error) {
	return (&roundRobinSwitcher{urls, 0}).GetProxy, nil
}*/