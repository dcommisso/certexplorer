package cmd

import (
	"github.com/dcommisso/certexplorer/certformatter"
)

type Configuration struct {
	certstore *certformatter.Certstore
}

func NewConfiguration() *Configuration {
	return &Configuration{
		certstore: certformatter.NewCertstore(),
	}
}
