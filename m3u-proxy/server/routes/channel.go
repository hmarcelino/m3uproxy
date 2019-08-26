package routes

import (
	"fmt"
	"github.com/hmarcelino/m3u-proxy/config"
	"net/http"
)

//
//import (
//	"fmt"
//	"github.com/hmarcelino/m3u-proxy/config"
//	"log"
//	"net/http"
//	"net/http/httputil"
//	"net/url"
//)
//
////type Proxy struct {
////	rawTarget     *url.URL
////	proxyHost     *httputil.ReverseProxy
////	requestString string
////}
//
func ChannelRoute(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	return "/channels/{Id}", func(w http.ResponseWriter, r *http.Request) {

	}
}
//
//func NewProxy(channelListAddr string) *Proxy {
//	u, err := url.Parse(channelListAddr)
//	if err != nil {
//		log.Fatalf("Error parsing m3u url %s, %v", channelListAddr, err)
//	}
//
//	uHost, _ := url.Parse(u.Scheme + u.Host)
//
//	proxy := httputil.NewSingleHostReverseProxy(uHost)
//	proxy.Director = func(request *http.Request) {
//
//	}
//
//	return &Proxy{
//		rawTarget:     u,
//		proxyHost:,
//		requestString: u.RequestURI(),
//	}
//}
//
//// The return should match the previous route pattern.
//// Http://host:port/channels/channelId
func NewChannelProxy(config *config.Config, id string) string {
	return fmt.Sprintf(
		"http://%s:%d/channels/%s",
		config.Server.Hostname,
		config.Server.Port,
		id,
	)
}
