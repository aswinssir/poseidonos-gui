// +build sslon

package handler

import (
	"A-module/log"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

var conn *tls.Conn

func Dial(network, addr string) (*tls.Conn, error) {

	caCert, err := ioutil.ReadFile("/etc/ibofos/cert/cert.crt")

	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	conf := &tls.Config{
		RootCAs: caCertPool,
	}

	conn, err = tls.Dial(network, addr, conf)

	return conn, err
}