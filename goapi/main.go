package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// https://github.com/yongs2/golang-exercise/blob/9cca75b2c0f22c303151380a4bfe67002ff9970d/12.golangbyexample/47.network/http_ip.go

func main() {
	// Methods 1
	GetsIP6Address()

}

// Method 1
func GetsIP6Address() {
	
	http.HandleFunc("/", ExampleHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ip, _ := getIP(r)
	resp, _ := json.Marshal(map[string]string{
		"ip": ip,
	})
	w.Write(resp)
}

//GetIP returns the IP address
func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		fmt.Printf("X-REAL-IP=[%v]\n", ip)
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			fmt.Printf("X-FORWARDED-FOR=[%v]\n", ip)
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Printf("SplitHostPort(%v).Err=[%v]\n", r.RemoteAddr, err)
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		fmt.Printf("RemoteAddr=[%v]\n", ip)
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}
