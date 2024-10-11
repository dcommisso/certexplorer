package cmd

import (
	"github.com/dcommisso/cabundleinspect/internal/certs"
	"github.com/spf13/pflag"
)

type Configuration struct {
	flags     *pflag.FlagSet
	certstore *certs.Certstore
}

func NewConfiguration() *Configuration {
	return &Configuration{
		certstore: certs.NewCertstore(),
	}
}

func (c *Configuration) LoadFlags(flags *pflag.FlagSet) {
	c.flags = flags
}
