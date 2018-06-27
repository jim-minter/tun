package main

import (
	"crypto/rsa"
	"crypto/x509"
	"io/ioutil"
	"net"
	"os"

	"github.com/jim-minter/azure-helm/pkg/tls"
)

func writeKey(filename string, key *rsa.PrivateKey) error {
	b, err := tls.PrivateKeyAsBytes(key)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, b, 0600)
}

func writeCert(filename string, cert *x509.Certificate) error {
	b, err := tls.CertAsBytes(cert)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, b, 0666)
}

func run() error {
	err := os.RemoveAll("certs")
	if err != nil {
		return err
	}

	err = os.MkdirAll("certs", 0777)
	if err != nil {
		return err
	}

	caKey, caCert, err := tls.NewCA("ca")
	if err != nil {
		return err
	}

	err = writeKey("certs/ca.key", caKey)
	if err != nil {
		return err
	}

	err = writeCert("certs/ca.crt", caCert)
	if err != nil {
		return err
	}

	serverKey, serverCert, err := tls.NewCert("server", nil, nil, []net.IP{net.ParseIP(os.Args[1])}, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, caKey, caCert)
	if err != nil {
		return err
	}

	err = writeKey("certs/server.key", serverKey)
	if err != nil {
		return err
	}

	err = writeCert("certs/server.crt", serverCert)
	if err != nil {
		return err
	}

	clientKey, clientCert, err := tls.NewCert("client", nil, nil, nil, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, caKey, caCert)
	if err != nil {
		return err
	}

	err = writeKey("certs/client.key", clientKey)
	if err != nil {
		return err
	}

	err = writeCert("certs/client.crt", clientCert)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
