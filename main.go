package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	h := func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("err:%v r.RemoteAddr:%s", err, r.RemoteAddr)
			return
		}
		log.Printf("ip:%s", ip)
		w.Write([]byte(ip))
	}
	http.HandleFunc("/ip", h)
	http.HandleFunc("/whatismyip", h)
	http.HandleFunc("/what-is-my-ip", h)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
