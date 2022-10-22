package artifacts

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifactTypeDetectorErrors(t *testing.T) {
	var sb strings.Builder

	t.Setenv("N3DR_MAVEN_UPLOAD_REGEX_VERSION", "notAMatch")
	err := artifactTypeDetector(&sb, "a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar", false)
	assert.EqualError(t, err, "check whether regexVersion: 'notAMatch' and regexClassifier: '(-(.*?(\\-([\\w.]+))?)?)?' match the artifact: 'a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar'")

	t.Setenv("N3DR_MAVEN_UPLOAD_REGEX_CLASSIFIER", "notAMatch2")
	err = artifactTypeDetector(&sb, "a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar", false)
	assert.EqualError(t, err, "check whether regexVersion: 'notAMatch' and regexClassifier: 'notAMatch2' match the artifact: 'a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar'")
}

func TestArtifactTypeDetector(t *testing.T) {
	var sb strings.Builder

	t.Setenv("N3DR_MAVEN_UPLOAD_REGEX_VERSION", `(([A-Za-z\d\-_]+)|(([a-z\d\.]+)))`)
	t.Setenv("N3DR_MAVEN_UPLOAD_REGEX_CLASSIFIER", `(-(.*?(\-([\w.]+))?)?)?`)

	if err := artifactTypeDetector(&sb, "3rdparty-maven-gradle-plugins/com/github/ben-manes/gradle-versions-plugin/0.30.0/gradle-versions-plugin-0.30.0.jar", false); err != nil {
		t.Error(err)
	}
	if err := artifactTypeDetector(&sb, "a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar", false); err != nil {
		t.Error(err)
	}
	if err := artifactTypeDetector(&sb, "d/e/f-e-d/4.5.6/d.e.f-4.5.6.pom", false); err != nil {
		t.Error(err)
	}
	if err := artifactTypeDetector(&sb, "d/e/f-e-d/7.8.9-2/d.e.f-7.8.9-2.war", false); err != nil {
		t.Error(err)
	}
	if err := artifactTypeDetector(&sb, "hello/w_o_r_l_d/1.0.1/w_o_r_l_d-1.0.1.jar", false); err != nil {
		t.Error(err)
	}
	if err := artifactTypeDetector(&sb, "hello/World/1.0/World-1.0.jar", false); err != nil {
		t.Error(err)
	}

	assert.Equal(t, "maven2.asset0=@3rdparty-maven-gradle-plugins/com/github/ben-manes/gradle-versions-plugin/0.30.0/gradle-versions-plugin-0.30.0.jar,maven2.asset0.extension=jar,maven2.asset1=@a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar,maven2.asset1.extension=jar,maven2.asset1.classifier=dataset,maven2.asset2=@d/e/f-e-d/4.5.6/d.e.f-4.5.6.pom,maven2.asset2.extension=pom,maven2.asset3=@d/e/f-e-d/7.8.9-2/d.e.f-7.8.9-2.war,maven2.asset3.extension=war,maven2.asset4=@hello/w_o_r_l_d/1.0.1/w_o_r_l_d-1.0.1.jar,maven2.asset4.extension=jar,maven2.asset5=@hello/World/1.0/World-1.0.jar,maven2.asset5.extension=jar,", sb.String())
}
