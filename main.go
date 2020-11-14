package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"lhosttls/utils"
	"net/http"
	"os"
)

func main() {

	server := getServer()
	http.HandleFunc("/", myHandler)
	server.ListenAndServeTLS("", "")

}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling request")
	fmt.Fprintln(w, "Hi there")
}

func getServer() *http.Server {

	data, err := ioutil.ReadFile("ca/minica.pem")

	if err != nil {
		fmt.Printf("error in reading the root certificate:%v\n", err)
		os.Exit(1)
	}

	cp := x509.NewCertPool()

	cp.AppendCertsFromPEM(data)

	tls := &tls.Config{
		ClientCAs:             cp,
		ClientAuth:            tls.RequireAndVerifyClientCert,
		GetCertificate:        utils.CertReqFunc("ca/ramesh-server/cert.pem", "ca/ramesh-server/key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tls,
	}

	return server

}
