package artifacts

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifactTypeDetector(t *testing.T) {
	var sb strings.Builder

	artifactTypeDetector(&sb, "a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar")
	artifactTypeDetector(&sb, "d/e/f-e-d/4.5.6/d.e.f-4.5.6.pom")
	artifactTypeDetector(&sb, "d/e/f-e-d/7.8.9-2-37zgb398/d.e.f-7.8.9-2-37zgb398.war")
	artifactTypeDetector(&sb, "d/e/f-e-d/9.8.7-81ae5835bb36126fe8091e82t14521841d8y0133/d.e.f-9.8.7-81ae5835bb36126fe8091e82t14521841d8y0133.war")

	assert.Equal(t, "maven2.asset0=@a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar,maven2.asset0.extension=jar,maven2.asset0.classifier=dataset,maven2.asset1=@d/e/f-e-d/7.8.9-2-37zgb398/d.e.f-7.8.9-2-37zgb398.war,maven2.asset1.extension=war,maven2.asset1.classifier=37zgb398,maven2.asset2=@d/e/f-e-d/9.8.7-81ae5835bb36126fe8091e82t14521841d8y0133/d.e.f-9.8.7-81ae5835bb36126fe8091e82t14521841d8y0133.war,maven2.asset2.extension=war,maven2.asset2.classifier=81ae5835bb36126fe8091e82t14521841d8y0133,", sb.String())
}
