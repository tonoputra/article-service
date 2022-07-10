package helper

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CertificateInterface interface {
	GetCustomTLSConfig() (*tls.Config, error)
	GetCertFile() string
}

type certificate struct {
}

func Certificate() CertificateInterface {
	return &certificate{}
}

func (c *certificate) GetCustomTLSConfig() (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	caFilePath := c.GetCertFile()

	certs, err := ioutil.ReadFile(caFilePath)
	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)
	if !ok {
		return tlsConfig, errors.New("failed parsing pem file")
	}

	return tlsConfig, nil
}

func (*certificate) GetCertFile() string {
	cwd, _ := os.Getwd()
	pemFile := "usr/bin/rds-combined-ca-bundle.pem"
	if os.Getenv("CA_FILE_PATH") != "" {
		pemFile = os.Getenv("CA_FILE_PATH")
	}
	return filepath.Join(cwd, pemFile)
}
