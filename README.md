# M3U Proxy

Proxy written in Go. It will process an M3U response and transform the all channel addresses with new addresses. 
It was built taking with the purpose that all traffic goes through it.

### Endpoints

* GET /
```bash
$> curl http://${host}:${port}

Welcome to m3u proxy
```

* GET /ping
```bash
$> curl http://${host}:${port}/ping

pong
```

* GET /channels
```
$> curl http://localhost:9090/ping

#EXTM3U
#EXTINF:-1 tvg-id="ABC FHD" tvg-name="ABC FHD" tvg-logo="http://....png" group-title="group A"
http://localhost:9090/channels/5287
#EXTINF:-1 tvg-id="ABC HD" tvg-name="ABC HD" tvg-logo="http://....png" group-title="group A"
http://localhost:9090/channels/984
...
```

* GET /channels/{id}
```
$> curl http://localhost:9090/channel/984

... stream ...
```

* GET /channels/info/{id}
```
$> curl http://localhost:9090/channel/info/984

{
    Id: "984",
    Source: {
        Scheme: "http",
        Opaque: "",
        User: null,
        Host: "<originalHost>:<originalPort>,
        Path: "<original request uri to channel>",
        RawPath: "",
        ForceQuery: false,
        RawQuery: "",
        Fragment: ""
    }
}
``` 

### Requirements:
* Golang >= 1.11
* docker (if you whish to build the container)   


### Setup
```bash
# Build the proxy
go build m3u-proxy/main.go

# Run locally
# Requires config file. Look for example in config/config-dev.yml
go run m3u-proxy/main.go -file <path to config file>

# Run locally with environment variables
# Useful when running in a docker container
export M3UPROXY_PORT="9090"
export M3UPROXY_HOSTNAME="localhost"
export M3UPROXY_M3U_URL="http://localhost:8080/channels.m3u"
go run m3u-proxy/main.go
```