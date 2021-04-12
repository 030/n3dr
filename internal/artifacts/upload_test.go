package artifacts

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifactTypeDetectorErrors(t *testing.T) {
	var sb strings.Builder

	os.Setenv("N3DR_MAVEN_UPLOAD_REGEX_VERSION", "notAMatch")
	err := artifactTypeDetector(&sb, "a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar", false)
	assert.EqualError(t, err, "check whether regexVersion: 'notAMatch' and regexClassifier: '(-(.*?(\\-([\\w.]+))?)?)?' match the artifact: 'a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar'")

	os.Setenv("N3DR_MAVEN_UPLOAD_REGEX_CLASSIFIER", "notAMatch2")
	err = artifactTypeDetector(&sb, "a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar", false)
	assert.EqualError(t, err, "check whether regexVersion: 'notAMatch' and regexClassifier: 'notAMatch2' match the artifact: 'a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar'")
}

func TestArtifactTypeDetector(t *testing.T) {
	var sb strings.Builder

	os.Setenv("N3DR_MAVEN_UPLOAD_REGEX_VERSION", `(([a-z\d\-]+)|(([a-z\d\.]+)))`)
	os.Setenv("N3DR_MAVEN_UPLOAD_REGEX_CLASSIFIER", `(-(.*?(\-([\w.]+))?)?)?`)

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

	assert.Equal(t, "maven2.asset0=@3rdparty-maven-gradle-plugins/com/github/ben-manes/gradle-versions-plugin/0.30.0/gradle-versions-plugin-0.30.0.jar,maven2.asset0.extension=jar,maven2.asset1=@a/b/c.b.a/1.2.3/a-b-c-1.2.3-dataset.jar,maven2.asset1.extension=jar,maven2.asset1.classifier=dataset,maven2.asset2=@d/e/f-e-d/4.5.6/d.e.f-4.5.6.pom,maven2.asset2.extension=pom,maven2.asset3=@d/e/f-e-d/7.8.9-2/d.e.f-7.8.9-2.war,maven2.asset3.extension=war,", sb.String())
}
