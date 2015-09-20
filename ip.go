package main

import (
	"errors"
	"fmt"
	flag "github.com/ogier/pflag"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {

	var printLocal *bool = flag.BoolP("local", "l", false, "Print local IP address")
	var printPublic *bool = flag.BoolP("public", "p", false, "Print public IP address")
	flag.Parse()

	ips := make(chan string, 2)

	go func() {
		ip, _ := localIP()
		ips <- ip
		ips <- externalIP()
	}()

	if !*printPublic && *printLocal {
		fmt.Println(<-ips)
		return
	}

	if *printPublic && !*printLocal {
		<-ips
		fmt.Println(<-ips)
		return
	}

	if *printLocal && *printPublic {
		fmt.Println(<-ips)
		fmt.Println(<-ips)
		return
	}

	fmt.Println("Local IP:", <-ips)
	fmt.Println("Public IP:", <-ips)
}

func externalIP() string {
	resp, err := http.Get("http://canihazip.com/s")
	if err != nil {
		return "-"
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body[:len(body)])
}

// https://code.google.com/p/whispering-gophers/source/browse/util/helper.go
func localIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
