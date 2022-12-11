package main

import (
	"fmt"
	"github.com/cczyWyc/crawler-geektime/collect"
	"github.com/cczyWyc/crawler-geektime/proxy"
	"time"
)

func main() {
	proxyURLs := []string{"http://127.0.0.1:7890"}
	p, err := proxy.RoundRobinProxySwitcher(proxyURLs...)
	if err != nil {
		fmt.Println("RoundRobinProxySwitcher failed")
	}
	url := "https://www.google.com"

	var f collect.FetCher = collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   p,
	}
	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("read contect faield: %v\n", err)
		return
	}
	fmt.Println(string(body))
}
