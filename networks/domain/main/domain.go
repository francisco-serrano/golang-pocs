package main

import (
	"fmt"
	"net"
)

func lookIP(addr string) ([]string, error) {
	hosts, err := net.LookupAddr(addr)
	if err != nil {
		return nil, err
	}

	return hosts, err
}

func lookHostname(hostname string) ([]string, error) {
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}

	return IPs, err
}

func exampleA() {
	//input := "127.0.0.1"
	//input := "172.20.16.0"
	input := "packtpub.com"

	IPaddress := net.ParseIP(input)

	if IPaddress == nil {
		IPs, err := lookHostname(input)
		if err == nil {
			for _, ip := range IPs {
				fmt.Println(ip)
			}
		}
	} else {
		hosts, err := lookIP(input)
		if err == nil {
			for _, host := range hosts {
				fmt.Println(host)
			}
		}
	}
}

func exampleB() {
	domain := "packtpub.com"
	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}

func exampleC() {
	domain := "packtpub.com"
	MXs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, MX := range MXs {
		fmt.Println(MX.Host)
	}
}

func main() {
	//exampleA()
	//exampleB()
	exampleC()
}
