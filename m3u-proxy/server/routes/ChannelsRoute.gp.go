package routes

import (
	"bytes"
	"github.com/hmarcelino/m3u-proxy/config"
	"github.com/hmarcelino/m3u-proxy/db"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

type Proxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func ChannelsRouter(config *config.Config) (string, func(w http.ResponseWriter, r *http.Request)) {
	proxyInstance := NewProxy(config.M3u.url)
	proxyInstance.proxy.ModifyResponse = func(resp *http.Response) error {
		db.ClearDB()

		b, err := ioutil.ReadAll(resp.Body) //Read html
		if err != nil {
			return err
		}
		err = resp.Body.Close()
		if err != nil {
			return err
		}

		payload := string(b)
		lines := strings.Split(payload, "\n")
		for _, line := range lines {

			// Line holding metadata. Something like
			// #EXTINF:-1 tvg-id="ABC HD" tvg-name="ABC FHD" ...
			if !strings.HasPrefix(line, "http://") {
				continue
			}

			channel, err := db.RegisterChannel(line)
			if err != nil {
				log.Fatalf("Error registering m3u url. %v", err)
			}

		}

		b = bytes.Replace(b, []byte("server"), []byte("schmerver"), -1) // replace html
		body := ioutil.NopCloser(bytes.NewReader(b))
		resp.Body = body
		resp.ContentLength = int64(len(b))
		resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
		return nil
	}

	return "/channels.m3u", func(w http.ResponseWriter, r *http.Request) {
		proxyInstance.proxy.ServeHTTP(w, r)
	}
}

func NewProxy(target string) *Proxy {
	targetUrl, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Error parsing m3u url %s, %v", target, err)
	}

	return &Proxy{
		target: targetUrl,
		proxy:  httputil.NewSingleHostReverseProxy(targetUrl),
	}
}
