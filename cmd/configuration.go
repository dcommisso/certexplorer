package cmd

import (
	"github.com/dcommisso/cabundleinspect/certformatter"
)

type Configuration struct {
	certstore *certformatter.Certstore
}

func NewConfiguration() *Configuration {
	return &Configuration{
		certstore: certformatter.NewCertstore(),
	}
}
