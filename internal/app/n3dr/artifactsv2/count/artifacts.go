package count

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/030/n3dr/internal/app/n3dr/artifactsv2/artifacts"
	"github.com/030/n3dr/internal/app/n3dr/connection"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/client/components"
	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
	log "github.com/sirupsen/logrus"
)

var (
	mu    sync.Mutex
	total int = 0
)

func (n *Nexus3) sort() error {
	input, err := os.Open(filepath.Clean(n.CsvFile + ".csv"))
	if err != nil {
		return err
	}
	defer func() {
		if err := input.Close(); err != nil {
			panic(err)
		}
	}()

	reader := csv.NewReader(input)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	pathColumn := 1 // Sort by second column
	sortByPath := func(i, j int) bool {
		return rows[i][pathColumn] < rows[j][pathColumn]
	}

	sort.Slice(rows, sortByPath)

	output, err := os.Create(filepath.Clean(n.CsvFile + "-sorted.csv"))
	if err != nil {
		return err
	}
	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	writer := csv.NewWriter(output)
	if err := writer.WriteAll(rows); err != nil {
		return err
	}
	writer.Flush()

	return nil
}

func (n *Nexus3) write(asset *models.AssetXO, repositoriesTotalArtifacts *int) {
	record := []string{asset.Repository, asset.Path, asset.DownloadURL, asset.Format, asset.ContentType, asset.LastDownloaded.String(), asset.LastModified.String()}
	checksums := []string{"md5", "sha1", "sha256", "sha512"}
	for _, checksum := range checksums {
		if value, ok := asset.Checksum[checksum]; ok {
			record = append(record, value.(string))
		}
	}

	if *repositoriesTotalArtifacts%1000 == 0 {
		log.Debugf("repository: '%s' counter: '%d'", asset.Repository, *repositoriesTotalArtifacts)
	}

	mu.Lock() // prevent "short write" issues by go routines that write on the same time by locking it
	if err := n.writer.Write(record); err != nil {
		panic(err)
	}
	n.writer.Flush()
	if err := n.writer.Error(); err != nil {
		panic(err)
	}
	mu.Unlock()
}

func (n *Nexus3) assets(assets []*models.AssetXO, repositoriesTotalArtifacts *int) {
	log.Tracef("assets: '%v'", assets)
	log.Debugf("number of assets: '%d'", len(assets))
	for _, asset := range assets {
		log.Tracef("asset: '%v'", asset)
		(*repositoriesTotalArtifacts)++
		total++

		if n.CsvFile != "" {
			n.write(asset, repositoriesTotalArtifacts)
		}
	}
}

func (n *Nexus3) items(items []*models.ComponentXO, repositoriesTotalArtifacts *int) {
	log.Debugf("number of items: '%d'", len(items))
	for _, item := range items {
		log.Tracef("item: '%v'", item)
		n.assets(item.Assets, repositoriesTotalArtifacts)
	}
}

func (n *Nexus3) artifact(continuationToken string, repositoriesTotalArtifacts *int, repo *models.AbstractAPIRepository) error {
	client, err := n.Nexus3.Client()
	if err != nil {
		return err
	}
	c := components.GetComponentsParams{ContinuationToken: &continuationToken, Repository: repo.Name}
	c.WithTimeout(time.Second * 60)
	resp, err := client.Components.GetComponents(&c)
	if err != nil {
		return fmt.Errorf("cannot get components: '%w'", err)
	}

	rgpl := resp.GetPayload()
	continuationToken = rgpl.ContinuationToken
	log.Tracef("continuationToken: '%s'", continuationToken)

	n.items(rgpl.Items, repositoriesTotalArtifacts)

	if continuationToken == "" {
		format := repo.Format

		for len(format) < 8 {
			format = format + " "
		}

		fmt.Printf("%d\t\t%s\t%s\t\t%s\n", *repositoriesTotalArtifacts, format, repo.Type, repo.Name)
		return nil
	}

	return n.artifact(continuationToken, repositoriesTotalArtifacts, repo)
}

func (n *Nexus3) csvWriter() (csvWriter, error) {
	f, err := os.Create(filepath.Clean(n.CsvFile + ".csv"))
	if err != nil {
		return csvWriter{}, err
	}

	w := csv.NewWriter(f)
	if err := w.Write([]string{"repo", "path", "downloadURL", "format", "contentType", "lastDownloaded", "lastModified", "checksumMd5", "checksumSha1", "checksumSha256", "checksumSha512"}); err != nil {
		return csvWriter{}, err
	}

	return csvWriter{file: f, writer: w}, err
}

func (n *Nexus3) Artifacts() error {
	c := connection.Nexus3{BasePathPrefix: n.BasePathPrefix, DockerHost: n.DockerHost, DockerPort: n.DockerPort, DockerPortSecure: n.DockerPortSecure, DownloadDirName: n.DownloadDirName, FQDN: n.FQDN, HTTPS: n.HTTPS, Pass: n.Pass, User: n.User}
	a := artifacts.Nexus3{Nexus3: &c}
	repos, err := a.Repos()
	if err != nil {
		return err
	}

	if n.CsvFile != "" {
		cw, err := n.csvWriter()
		if err != nil {
			return err
		}
		defer func() {
			if err := cw.file.Close(); err != nil {
				panic(err)
			}
		}()
		n.writer = cw.writer
	}

	var wg sync.WaitGroup
	fmt.Printf("COUNT\t\tFORMAT\t\tTYPE\t\tNAME\n")
	for _, repo := range repos {
		wg.Add(1)
		go func(repoPreventDataRace *models.AbstractAPIRepository) {
			defer wg.Done()

			repositoriesTotalArtifacts := 0
			log.Debugf("repositoriesTotalArtifacts: '%d'", repositoriesTotalArtifacts)

			if err := n.artifact("", &repositoriesTotalArtifacts, repoPreventDataRace); err != nil {
				panic(err)
			}
		}(repo)
	}
	wg.Wait()

	if n.Sort {
		if err := n.sort(); err != nil {
			panic(err)
		}
	}

	return nil
}
