package certs

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/dcommisso/cabundleinspect/internal/format"
)

const (
	CertificateBlockType = "CERTIFICATE"
)

type Certstore struct {
	certs map[int]Certificate
}

type Certificate struct {
	DecodedCertificate *x509.Certificate
	Source             string
}

func NewCertstore() *Certstore {
	return &Certstore{
		certs: make(map[int]Certificate),
	}
}

func (c *Certstore) Load(rawCerts []byte, source string) {
	nextFreeIndex := len(c.certs)
	for i := nextFreeIndex; len(rawCerts) > 0; i++ {
		var block *pem.Block
		block, rawCerts = pem.Decode(rawCerts)

		if block == nil {
			break
		}

		if block.Type != CertificateBlockType || len(block.Headers) != 0 {
			continue
		}

		cert, _ := x509.ParseCertificate(block.Bytes)
		c.certs[i] = Certificate{
			DecodedCertificate: cert,
			Source:             source,
		}
	}
}

func (c Certificate) GetSerialNumber() string {
	return format.ToColonNotation(c.DecodedCertificate.SerialNumber.Bytes())
}

func (c Certificate) GetIssuer() string {
	return c.DecodedCertificate.Issuer.String()
}

func (c Certificate) GetSubject() string {
	return c.DecodedCertificate.Subject.String()
}

func (c Certificate) GetNotBefore() string {
	return c.DecodedCertificate.NotBefore.String()
}

func (c Certificate) GetNotAfter() string {
	return c.DecodedCertificate.NotAfter.String()
}

func (c Certificate) GetSKID() string {
	return strings.ToUpper(format.ToColonNotation(c.DecodedCertificate.SubjectKeyId))
}

func (c Certificate) GetAKID() string {
	return strings.ToUpper(format.ToColonNotation(c.DecodedCertificate.AuthorityKeyId))
}

func (c Certificate) GetSANs() string {
	sans := []string{}

	for _, dns := range c.DecodedCertificate.DNSNames {
		sans = append(sans, fmt.Sprintf("DNS:%s", dns))
	}

	for _, ip := range c.DecodedCertificate.IPAddresses {
		sans = append(sans, fmt.Sprintf("IP Address:%s", ip.String()))
	}

	return strings.Join(sans, ", ")
}
