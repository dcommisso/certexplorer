package cmd

import (
	"github.com/dcommisso/cabundleinspect/internal/certs"
)

type Configuration struct {
	certstore *certs.Certstore
}

func NewConfiguration() *Configuration {
	return &Configuration{
		certstore: certs.NewCertstore(),
	}
}
