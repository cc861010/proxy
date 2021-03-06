package main

import (
        "flag"
        "log"
        "net/http"
        "fmt"

        "github.com/elazarl/goproxy"
        "golang.org/x/net/proxy"
)

/**
➜       ./proxy -h
        Usage of ./proxy:
          -p            proxy listen port (default "8080")
          -socks5       socks5 proxy (default "127.0.0.1:1080")
          -v            verbose

        test:
        http_proxy=http://localhost:8008 wget www.google.com
 */
func main() {
        PROXY_ADDR := flag.String("socks5", "127.0.0.1:8888", "socks5 proxy `address`")
        verbose := flag.Bool("v", true, "verbose")
        addr := flag.String("p", "8008", "proxy listen `port`")
        flag.Parse()

        // create a socks5 dialer
        dialer, err := proxy.SOCKS5("tcp", *PROXY_ADDR, nil, proxy.Direct)
        if err != nil {
                log.Fatalf("can't connect to the proxy: %v", err)
                return
        }else{
                fmt.Println("export HTTP_PROXY=http://localhost:8008")
                fmt.Println("export HTTPS_PROXY=http://localhost:8008")
        }

        proxy := goproxy.NewProxyHttpServer()
        proxy.Verbose = *verbose
        proxy.Tr.Dial = dialer.Dial
        log.Fatal(http.ListenAndServe(":"+*addr, proxy))
}


