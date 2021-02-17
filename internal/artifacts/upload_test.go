package cli

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifactTypeDetector(t *testing.T) {
	var sb strings.Builder

	artifactTypeDetector(&sb, "a/b/c.b.a/1.2.3/a-b-c-1.2.3.jar")
	artifactTypeDetector(&sb, "d/e/f-e-d/4.5.6/d.e.f-4.5.6.pom")

	assert.Equal(t, "maven2.asset0=@a/b/c.b.a/1.2.3/a-b-c-1.2.3.jar,maven2.asset0.extension=jar,maven2.asset1=@d/e/f-e-d/4.5.6/d.e.f-4.5.6.pom,maven2.asset1.extension=pom,", sb.String())
}
