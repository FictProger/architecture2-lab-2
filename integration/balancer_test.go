package integration

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

const baseAddress = "http://balancer:8090"

var client = http.Client{
	Timeout: 3 * time.Second,
}

func Test(t *testing.T) {
	time.Sleep(11 * time.Second)
	TestingT(t)
}

type IntegrationSuite struct{}

var _ = Suite(&IntegrationSuite{})

func (s *IntegrationSuite) TestBalancer(c *C) {
	var server string
	for i := 0; i < 10; i++ {
		resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", baseAddress))
		c.Assert(err, IsNil)

		c.Assert(resp.StatusCode, Equals, http.StatusOK)

		from := resp.Header.Get("lb-from")
		if server == "" {
			server = from
		} else {
			c.Assert(server, Equals, from)
		}
	}
}
