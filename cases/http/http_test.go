package http

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	caCert, err := ioutil.ReadFile("/opt/server/oss-manager/conf/manager.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	//html, err := client.Get("https://localhost:8000/api/v1/status")
	html, err := client.Get("http://www.jd.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	website, err := ioutil.ReadAll(html.Body)
	if err != nil {
		log.Fatal(err)
	}
	html.Body.Close()
	log.Println(string(website))
}
