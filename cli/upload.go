package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var s strings.Builder

func (n Nexus3) multipart(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		if filepath.Ext(path) == ".pom" {
			s.WriteString("maven2.asset1=@" + path + ",")
			s.WriteString("maven2.asset1.extension=pom,")
		}

		if filepath.Ext(path) == ".jar" {
			s.WriteString("maven2.asset2=@" + path + ",")
			s.WriteString("maven2.asset2.extension=jar,")
		}

		if filepath.Ext(path) == "sources.jar" {
			s.WriteString("maven2.asset3=@" + path + ",")
			s.WriteString("maven2.asset3.extension=sources-jar,")
		}
	}

	// url := n.URL + "/service/rest/v1/components?repository=" + n.Repository
	// u := mp.Upload{URL: url, Username: n.User, Password: n.Pass}
	// err2 := u.MultipartUpload("maven2.asset1=@" + path + ".pom,maven2.asset1.extension=pom,maven2.asset2=@" + path + ".jar,maven2.asset2.extension=jar")
	// if err2 != nil {
	// 	return err2
	// }

	return nil
}

// Upload posts an artifact as a multipart to a specific nexus3 repository
func (n Nexus3) Upload() error {
	// path := filepath.Join("testFiles", "file1")
	// url := n.URL + "/service/rest/v1/components?repository=" + n.Repository
	// u := mp.Upload{URL: url, Username: n.User, Password: n.Pass}
	// err := u.MultipartUpload("maven2.asset1=@" + path + ".pom,maven2.asset1.extension=pom,maven2.asset2=@" + path + ".jar,maven2.asset2.extension=jar")
	// if err != nil {
	// 	return err
	// }

	// files, err := ioutil.ReadDir(n.Repository)
	// if err != nil {
	// 	return err
	// }

	// for _, f := range files {
	// 	fmt.Println(f)
	// }

	err := filepath.Walk(n.Repository, n.multipart)
	if err != nil {
		return err
	}

	fmt.Println(strings.TrimSuffix(s.String(), ","))

	return nil
}
