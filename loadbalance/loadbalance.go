package loadbalance

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

// Backend holds the information about a backend server
type Backend struct {
	URL          *url.URL
	ReverseProxy *httputil.ReverseProxy
}

// LoadBalancer holds the list of backends and the current index
type LoadBalancer struct {
	backends []*Backend
	current  uint32
}

// NewLoadBalancer initializes a new LoadBalancer
func NewLoadBalancer(backends []*Backend) *LoadBalancer {
	return &LoadBalancer{
		backends: backends,
	}
}

// NextBackend returns the next backend in round-robin order
func (lb *LoadBalancer) NextBackend() *Backend {
	idx := atomic.AddUint32(&lb.current, 1)
	return lb.backends[(idx-1)%uint32(len(lb.backends))]
}

// ProxyHandler handles the incoming requests and forwards them to the backends
func (lb *LoadBalancer) ProxyHandler(c echo.Context) error {
	backend := lb.NextBackend()
	fmt.Println("Serve HTTP using", backend.URL)
	backend.ReverseProxy.ServeHTTP(c.Response(), c.Request())
	return nil
}

func LoadBalance() {
	e := echo.New()

	backend1URL, _ := url.Parse("http://localhost:8081")
	backend2URL, _ := url.Parse("http://localhost:8082")

	backends := []*Backend{
		{URL: backend1URL, ReverseProxy: httputil.NewSingleHostReverseProxy(backend1URL)},
		{URL: backend2URL, ReverseProxy: httputil.NewSingleHostReverseProxy(backend2URL)},
	}

	lb := NewLoadBalancer(backends)

	e.Any("/*", lb.ProxyHandler)

	e.Start(":8080")
}
