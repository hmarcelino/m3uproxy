# M3U Proxy

Proxy written in Go. It will process an M3U response and transform the all channel addresses with new addresses. 
It was built taking with the purpose that all traffic goes through it.

## Endpoints

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
```bash
$> curl http://localhost:9090/ping

#EXTM3U
#EXTINF:-1 tvg-id="ABC FHD" tvg-name="ABC FHD" tvg-logo="http://....png" group-title="group A"
http://localhost:9090/channels/5287
#EXTINF:-1 tvg-id="ABC HD" tvg-name="ABC HD" tvg-logo="http://....png" group-title="group A"
http://localhost:9090/channels/984
...
```

* GET /channels/{id}
```bash
$> curl http://localhost:9090/channel/984

... stream ...
```

* GET /channels/info/{id}
```bash
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

## FAQ

### Requirements:
* Golang >= 1.11
* docker (if you whish to build the container)

**How to build the proxy**
```bash
go build m3u-proxy/main.go
```

**How to run locally with config file**
Requires config file. Look for example in config/config-dev.yml
```bash
go run m3u-proxy/main.go -file <path to config file>
```

**How to run locally with environment variables**
This is useful when running in a docker container
```bash
export M3U_PROXY_PORT="9090"
export M3U_PROXY_HOSTNAME="localhost"
export M3U_PROXY_CHANNELS_URL="http://ilovetv.pt:8000/get.php?username=BjtVfUnp6b&password=tPMYWC11pg&type=m3u_plus&output=ts"
go run m3u-proxy/main.go

#or 

docker run -e M3U_PROXY_CHANNELS_URL="<valid url to m3u list>" -p 9090:9090  m3u-proxy:latest
 
```