package cmd

import (
	"github.com/dcommisso/certexplorer/certformatter"
)

type Configuration struct {
	certstore *certformatter.Certstore
	testMode  bool
}

func NewConfiguration() *Configuration {
	return &Configuration{
		certstore: certformatter.NewCertstore(),
		testMode:  false,
	}
}

func (c *Configuration) setTestMode(testMode bool) {
	c.testMode = testMode
}
