package cli

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

func (n *Nexus3) Bladibla(errs chan error, url string) error {
	resp, err := grequests.Get(url, &grequests.RequestOptions{Auth: []string{n.User, n.Pass}})
	if err != nil {
		return err
	}

	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		fmt.Printf("StatusCode not OK, but: '%v'", statusCode)
	}

	r := strings.NewReader(resp.String())
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	bodies := cascadia.MustCompile("tr td a").MatchAll(doc)
	for _, body := range bodies {
		time.Sleep(100 * time.Millisecond)
		go func(errs chan error, n *Nexus3, body *html.Node, url string) {
			log.Debugf("Go Channel length (inside go routine): '%d'", len(errs))
			errs <- n.wat(errs, body, url)
		}(errs, n, body, url)
	}
	for range bodies {
		if err := <-errs; err != nil {
			return err
		}
	}
	return nil
}

func (n *Nexus3) wat(errs chan error, body *html.Node, url string) error {
	ssss := goquery.NewDocumentFromNode(body).Text()

	if ssss != "Parent Directory" {
		log.Debug(ssss)
		url2 := url + "/" + ssss
		fmt.Println("URL: ", url2)
		fmt.Println("Extension: ", filepath.Ext(url2))

		if filepath.Ext(url2) == ".tgz" {
			re, err := regexp.Compile("^(.*)/service\\/rest\\/repository\\/browse\\/(.*)\\/(.*)$")
			if err != nil {
				return err
			}
			if !re.MatchString(url2) {
				return fmt.Errorf("No MATCH!!!!!!!!!!: %v", url2)
			}
			group := re.FindStringSubmatch(url2)
			url2 = group[1] + "/repository/" + group[2] + "/-/" + group[3]

			fmt.Println("Download URL: " + url2)
			resp, _ := grequests.Get(url2, &grequests.RequestOptions{Auth: []string{n.User, n.Pass}})
			fmt.Println("FILEPATH", filepath.Join("testi/", group[2], group[3]))
			os.MkdirAll(filepath.Join("testi/", group[2]), os.ModePerm)
			if err := resp.DownloadToFile(filepath.Join("./testi/", group[2], group[3])); err != nil {
				return err
			}
		}

		if err := n.Bladibla(errs, url2); err != nil {
			return err
		}
	}
	return nil
}
