package cli

const (
	pingURL = "http://localhost:8081/service/metrics/ping"
)

func Sum(a int, b int) int {
	return a + b
}
