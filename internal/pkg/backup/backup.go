package backup

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type ContinuationToken string

func ContinuationTokenInJSON(json string) (ContinuationToken, error) {
	if value := gjson.Get(json, "continuationToken"); !value.Exists() {
		log.Debugf("JSON: '%s'", json)
		return "", fmt.Errorf("continuationToken does not exist in json")
	} else {
		return ContinuationToken(value.String()), nil
	}
}
