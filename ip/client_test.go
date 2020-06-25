package ip

import (
	"testing"

	"github.com/aerialls/scaleway-ddns/tests"

	"github.com/stretchr/testify/assert"
)

func TestClientValidIPv4(t *testing.T) {
	ts := tests.StartWebServer(map[string]string{
		"/ipv4": "8.8.8.8",
	})
	defer ts.Close()

	ip, err := GetPublicIP(ts.URL + "/ipv4")

	assert.Nil(t, err)
	assert.Equal(t, ip, "8.8.8.8")
}

func TestClientValidIPv4WithPublicService(t *testing.T) {
	ip, err := GetPublicIP("https://api.ipify.org")

	assert.Nil(t, err)
	assert.NotEmpty(t, ip)
}

func TestClientValidIPv4WithReturnLine(t *testing.T) {
	ts := tests.StartWebServer(map[string]string{
		"/ipv4": "8.8.8.8\n\n",
	})
	defer ts.Close()

	ip, err := GetPublicIP(ts.URL + "/ipv4")

	assert.Nil(t, err)
	assert.Equal(t, ip, "8.8.8.8")
}

func TestClientValidIPv6(t *testing.T) {
	ts := tests.StartWebServer(map[string]string{
		"/ipv6": "2a00:1450:4007:816::2003",
	})
	defer ts.Close()

	ip, err := GetPublicIP(ts.URL + "/ipv6")

	assert.Nil(t, err)
	assert.Equal(t, ip, "2a00:1450:4007:816::2003")
}
