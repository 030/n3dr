package count

import (
	"encoding/csv"
	"os"

	"github.com/030/n3dr/internal/app/n3dr/connection"
)

type csvWriter struct {
	file   *os.File
	writer *csv.Writer
}

type Nexus3 struct {
	*connection.Nexus3
	CsvFile string
	writer  *csv.Writer
}
