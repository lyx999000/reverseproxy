package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("https://history.nasa.gov/SP-4206")
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	normalPage := func(w http.ResponseWriter, req *http.Request) {

		req.Host = req.URL.Host
		if req.RequestURI == "/" && req.URL.Path == "/" {
			req.RequestURI = "/contents.htm"
			req.URL.Path = "/contents.htm"
		}
		proxy.ServeHTTP(w, req)

	}
	warningPage := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "\t\t********* WARNING ********* \n")
		io.WriteString(w, "\tHey! YOU! \n")
		io.WriteString(w, "\tPerpetrator From IP: "+req.RemoteAddr+"\n")
		io.WriteString(w, "\tIt's the Command's decision to withhold the matters of Chapter 8 from public concern \n")
		io.WriteString(w, "\tTurn back! \n")
		io.WriteString(w, "\t\t********* WARNING ********* \n")
	}

	http.HandleFunc("/", normalPage)
	http.HandleFunc("/ch8.htm", warningPage)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
