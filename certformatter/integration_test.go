package certformatter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	filesToLoad := []string{"testdata/fewCertificates.pem", "testdata/example.com.crt"}

	cs := NewCertstore()
	csformatter := cs.NewFormatter()

	for _, fname := range filesToLoad {
		b, err := os.ReadFile(fname)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		cs.Load(b, fname)
	}

	certsToRender := []FormattedCertificate{}
	certsToRender = append(certsToRender, csformatter.GetFormattedCertificate(0))
	certsToRender = append(certsToRender, csformatter.GetFormattedCertificate(4))

	gotRendered, err := csformatter.ComposeFormattedCertificates(certsToRender, []Outputfield{OutputFieldSubject, OutputFieldValidity, OutputFieldSourceFile})
	if err != nil {
		t.Errorf("error: %v", err)
	}

	expectedRendered := `[0] Subject: OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES
    Validity
        Not Before: 2008-10-29 15:59:56 +0000 UTC
        Not After : 2030-01-01 00:00:00 +0000 UTC
    From file: testdata/fewCertificates.pem

[4] Subject: CN=example.com
    Validity
        Not Before: 2024-10-07 15:44:12 +0000 UTC
        Not After : 2034-10-05 15:44:12 +0000 UTC
    From file: testdata/example.com.crt`

	assert.Equal(t, expectedRendered, gotRendered)
}
