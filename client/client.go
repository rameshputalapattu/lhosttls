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

	client := getClient()

	resp, err := client.Get("https://ramesh-server:8080")

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Status: %s Body:  %s\n", resp.Status, string(body))

}

func getClient() *http.Client {
	data, err := ioutil.ReadFile("../ca/minica.pem")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	//cp, err := x509.SystemCertPool()
	cp := x509.NewCertPool()
	if err != nil {
		fmt.Printf("error getting cp:%v\n", err)
		os.Exit(1)
	}

	cp.AppendCertsFromPEM(data)
	config := &tls.Config{

		RootCAs:               cp,
		GetClientCertificate:  utils.ClientCertReqFunc("../ca/gophers/cert.pem", "../ca/gophers/key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}
