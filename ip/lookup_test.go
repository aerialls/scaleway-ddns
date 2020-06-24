package ip

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoServicesShouldReturnSameIPv4(t *testing.T) {
	aIP, err := GetPublicIP("https://api.ipify.org")
	assert.Nil(t, err)

	bIP, err := GetPublicIP("https://api-ipv4.ip.sb/ip")
	assert.Nil(t, err)

	assert.Equal(t, aIP, bIP)
}

func TestTwoServicesShouldReturnSameIPv6(t *testing.T) {
	aIP, err := GetPublicIP("https://v6.ident.me/")
	if err != nil && strings.Contains(err.Error(), "connect: no route to host") {
		t.Skip("IPv6 connectivity not available")
	}

	assert.Nil(t, err)

	bIP, err := GetPublicIP("https://api-ipv6.ip.sb/ip")
	assert.Nil(t, err)

	assert.Equal(t, aIP, bIP)
}
