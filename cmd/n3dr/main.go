package main

import (
	"github.com/030/n3dr/internal/app/n3dr/logger"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	if err := logger.File("cmd/n3dr/main", log); err != nil {
		panic(err)
	}
}

func main() {
	execute()
}
