package artifacts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDirHome = "/tmp/n3drtest"
)

func TestValidateNexusURLError(t *testing.T) {
	nexus3URLs := []string{
		"boo.foo",
		"http2://boo.foo",
		"http://boo.foo/",
		"https://boo.foo/",
		"hts://10.20.30.40",
		"https://boo-foo/",
		"https://boo-foo.foo:8080/",
	}
	for _, nexus3URL := range nexus3URLs {
		n := Nexus3{URL: nexus3URL}
		assert.EqualError(t, n.ValidateNexusURL(), "the Nexus3 URL seems to be incorrect. Verify that it complies to the regex that is defined in the 'Nexus3 Struct' and that it does not end with a '/'. Error: 'URL: regular expression mismatch'")
	}
}

func TestValidateNexusURL(t *testing.T) {
	nexus3URLs := []string{
		"http://boo.foo",
		"http://boo.foo-boo-foo",
		"http://1.2.3.4",
		"http://boo.foo-boo-foo:8765",
		"https://boo.foo",
		"https://boo.foo-boo-foo",
		"https://boo.foo-boo-foo:1234",
		"https://8.9.2.33:1234",
	}
	for _, nexus3URL := range nexus3URLs {
		n := Nexus3{URL: nexus3URL}
		err := n.ValidateNexusURL()
		if err != nil {
			t.Error(err)
		}
		assert.Nil(t, err, n.URL)
	}
}
