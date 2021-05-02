// package backup

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"path"
// 	"path/filepath"
// 	"regexp"
// 	"strings"
// 	"time"

// 	"github.com/PuerkitoBio/goquery"
// 	"github.com/andybalholm/cascadia"
// 	"github.com/levigross/grequests"
// 	log "github.com/sirupsen/logrus"
// 	"golang.org/x/net/html"
// )

// const nexus3RepoBrowseURI = "/service/rest/repository/browse/"

// type Nexus3 struct {
// 	BaseDir, Endpoint, Password, Repository, Regex, Username string
// }

// type nexusArtifactLabel string

// var toBeDownloadedArtifacts = 0

// func (n *Nexus3) AllArtifacts() error {
// 	log.Infof("Getting all artifacts from Nexus3 repository: '%s'...", n.Repository)

// 	errs := make(chan error)
// 	url := n.Endpoint + nexus3RepoBrowseURI + n.Repository
// 	if err := n.repositoryDirectoryOrArtifact(errs, url); err != nil {
// 		return err
// 	}

// 	time.Sleep(5 * time.Second)
// 	for {
// 		time.Sleep(500 * time.Millisecond)
// 		log.Infof("toBeDownloadedArtifacts: '%d' from nexus repository: '%s'", toBeDownloadedArtifacts, n.Repository)
// 		if toBeDownloadedArtifacts > 0 {
// 			if err := <-errs; err != nil {
// 				return err
// 			}
// 			toBeDownloadedArtifacts--
// 		}

// 		if toBeDownloadedArtifacts == 0 {
// 			break
// 		}
// 	}
// 	return nil
// }

// func (n *Nexus3) repositoryDirectoryOrArtifact(errs chan error, url string) error {
// 	log.Infof("Getting all raw HTML of nexusRepository: '%s'...", n.Repository)
// 	resp, err := n.repositoryRawHTML(url)
// 	if err != nil {
// 		return err
// 	}

// 	log.Infof("Getting all artifacts and repositories from nexus repository: '%s' in HTML...", n.Repository)
// 	directoriesAndArtifactsHtmlNode, err := repositoryRawHTMLDirectoriesOrArtifacts(resp.String())
// 	if err != nil {
// 		return err
// 	}

// 	for _, directoryOrArtifactHtmlNode := range directoriesAndArtifactsHtmlNode {
// 		directoryOrArtifact := goquery.NewDocumentFromNode(directoryOrArtifactHtmlNode).Text()
// 		log.Infof("directoryOrArtifact: '%s'", directoryOrArtifact)
// 		if err := n.downloadArtifactOrContinueSearchingInDirectory(errs, directoryOrArtifact, url); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // repositoryRawHTML returns the raw HTML that contains an overview of all
// // artifact directories that reside in a Nexus3 repository, e.g. nuget-hosted
// func (n *Nexus3) repositoryRawHTML(url string) (*grequests.Response, error) {
// 	resp, err := grequests.Get(url, &grequests.RequestOptions{Auth: []string{n.Username, n.Password}})
// 	if err != nil {
// 		return nil, err
// 	}

// 	statusCode := resp.StatusCode
// 	log.Debugf("URL: '%v'. StatusCode: '%v'. Response: '%s'", url, statusCode, resp.String())
// 	if statusCode != http.StatusOK {
// 		return nil, fmt.Errorf("StatusCode URL: '%s' not OK, but: '%d'. Enable debug mode to get the response", url, statusCode)
// 	}

// 	return resp, nil
// }

// // repositoryRawHTMLDirectoriesAndFiles returns the directories and files that
// // reside in a Nexus3 repository
// func repositoryRawHTMLDirectoriesOrArtifacts(s string) ([]*html.Node, error) {
// 	r := strings.NewReader(s)
// 	doc, err := html.Parse(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	directoriesAndArtifacts := cascadia.MustCompile("tr td a").MatchAll(doc)
// 	directoriesAndArtifactsSize := len(directoriesAndArtifacts)
// 	if directoriesAndArtifactsSize == 0 {
// 		return nil, fmt.Errorf("did not find any directories or artifacts; directoriesAndArtifactsSize: '%d'", directoriesAndArtifactsSize)
// 	}
// 	log.Infof("directoriesAndArtifactsSize: '%d'", directoriesAndArtifactsSize)

// 	return directoriesAndArtifacts, nil
// }

// // // urlRepoAndDir determines the repository name, e.g. nuget-hosted and artifact
// // // directory, e.g. @some-dir/some-artifact-dir
// // func urlRepoAndDir(url string) string {
// // 	re := regexp.MustCompile(`^.*/repository/(.*)/-/.*$`)
// // 	groups := re.FindStringSubmatch(url)
// // 	return groups[1]
// // }

// // download downloads an artifact, e.g. some-file.tgz from Nexus3
// func (n *Nexus3) download(url string) (nexusArtifactLabel, error) {
// 	resp, err := n.repositoryRawHTML(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	filePathDir := filepath.Join(n.BaseDir, n.Repository)
// 	filePath := filepath.Join(filePathDir, path.Base(url))

// 	log.Debugf("artifact filePathDir: '%s' and filePath: '%s'", filePathDir, filePath)
// 	if err := os.MkdirAll(filePathDir, os.ModePerm); err != nil {
// 		return "", err
// 	}

// 	if filepath.Ext(filePath) == ".nupkg" {
// 		filePath = filePath + "/n3dr-5.2.6.nupkg"
// 	}

// 	label, err := n.downloadAndPrintLabel(filePath, url, resp)
// 	if err != nil {
// 		return "", err
// 	}

// 	return label, nil
// }

// // downloadAndPrintLabel downloads an artifact and prints a label to indicate
// // what type of artifact has been downloaded, e.g. npm
// func (n *Nexus3) downloadAndPrintLabel(filePath, url string, resp *grequests.Response) (nexusArtifactLabel, error) {
// 	r, err := regexp.Compile(n.Regex)
// 	if err != nil {
// 		return "", err
// 	}

// 	var label nexusArtifactLabel
// 	if r.MatchString(url) {
// 		log.Infof("Downloading file: '%v'", url)
// 		if err := resp.DownloadToFile(filePath); err != nil {
// 			return "", err
// 		}
// 		log.Infof("File: '%s' downloaded. Total toBeDownloadedArtifacts: '%d'", filePath, toBeDownloadedArtifacts)

// 		switch fileExtension := filepath.Ext(filePath); fileExtension {
// 		case ".tgz":
// 			label = "*"
// 		default:
// 			log.Warningf("Unknown extension: '%v'", fileExtension)
// 			label = " unknownFileExtension "
// 		}
// 	} else {
// 		log.Warningf("Download of: '%s' skipped as it does not match with the regex: '%s'", filePath, n.Regex)
// 	}

// 	return label, nil
// }

// func (n *Nexus3) downloadArtifactOrContinueSearchingInDirectory(errs chan error, directoryOrArtifact, url string) error {
// 	log.Debugf("directoryOrArtifact: '%s'. URL: '%s'", directoryOrArtifact, url)

// 	if directoryOrArtifact != "Parent Directory" {
// 		url = url + "/" + directoryOrArtifact
// 		ext := filepath.Ext(url)
// 		log.Infof("URL: '%s'. Extension: '%s'", url, ext)

// 		if ext == ".tgz" {
// 			go func() {
// 				toBeDownloadedArtifacts++
// 				log.Infof("Downloading: '%s'", url)
// 				label, err := n.download(downloadURL(url))
// 				fmt.Print(label)
// 				errs <- err
// 			}()
// 			return nil
// 		}
// 		// if ext == ".nupkg" {
// 		// 	go func() {
// 		// 		toBeDownloadedArtifacts++
// 		// 		log.Infof("Downloading: '%s'", url)
// 		// 		label, err := n.download(downloadURLNupkg(url))
// 		// 		fmt.Print(label)
// 		// 		errs <- err
// 		// 	}()
// 		// 	return nil
// 		// }

// 		if err := n.repositoryDirectoryOrArtifact(errs, url); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
