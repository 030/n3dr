package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var foldersWithPOM strings.Builder

func (n Nexus3) detectFoldersWithPOM(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && filepath.Ext(path) == ".pom" {
		foldersWithPOM.WriteString(filepath.Dir(path) + ",")
	}
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

	// err := filepath.Walk(n.Repository, n.multipart)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(strings.TrimSuffix(s.String(), ","))

	err2 := filepath.Walk(n.Repository, n.detectFoldersWithPOM)
	if err2 != nil {
		return err2
	}

	foldersWithPOMString := strings.TrimSuffix(foldersWithPOM.String(), ",")
	foldersWithPOMStringSlice := strings.Split(foldersWithPOMString, ",")

	for _, v := range foldersWithPOMStringSlice {
		var s strings.Builder
		err := filepath.Walk(v, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !f.IsDir() {
				if filepath.Ext(path) == ".pom" {
					s.WriteString("maven2.asset1=@" + path + ",")
					s.WriteString("maven2.asset1.extension=pom,")
				}

				if filepath.Ext(path) == ".jar" {
					s.WriteString("maven2.asset2=@" + path + ",")
					s.WriteString("maven2.asset2.extension=jar,")
				}

				if filepath.Ext(path) == ".sources-jar" {
					s.WriteString("maven2.asset3=@" + path + ",")
					s.WriteString("maven2.asset3.extension=sources-jar,")
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		multipartString := strings.TrimSuffix(s.String(), ",")
		fmt.Println(multipartString)
		// url := n.URL + "/service/rest/v1/components?repository=" + n.Repository
		// u := mp.Upload{URL: url, Username: n.User, Password: n.Pass}
		// err2 := u.MultipartUpload(multipartString)
		// if err2 != nil {
		// 	return err
		// }
	}
	return nil
}
