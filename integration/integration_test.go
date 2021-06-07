package integration

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func BenchmarkBalancer(b *testing.B) {
	var client = http.Client{
		Timeout: 3 * time.Second,
	}
	const baseAddress = "http://balancer:8090"

	for i := 0; i < b.N; i++ {
		resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress))
		if err != nil {
			b.Error(err)
		}
		if resp.StatusCode != http.StatusOK {
			b.Error(fmt.Errorf("status code not OK"))
		}
	}
}
