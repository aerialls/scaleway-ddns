package ip

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// GetPublicIP returns the current public IP (IPv4 or IPv6 depending of the URL)
func GetPublicIP(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unable to get public IP from endpoint %s", url)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	rawIP := strings.TrimSpace(string(body))
	ip := net.ParseIP(rawIP)

	if ip == nil {
		return "", errors.New("the response is not a valid IP address")
	}

	return ip.String(), nil
}
