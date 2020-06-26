# scaleway-ddns

[![Build status](https://github.com/aerialls/scaleway-ddns/workflows/Test/badge.svg)](https://github.com/aerialls/scaleway-ddns/actions?query=workflow%3ATest)
[![Go report card](https://goreportcard.com/badge/github.com/aerialls/scaleway-ddns)](https://goreportcard.com/report/github.com/aerialls/scaleway-ddns)
![Go version](https://img.shields.io/github/go-mod/go-version/aerialls/scaleway-ddns)

Dynamic DNS service based on [Scaleway DNS](https://console.scaleway.com/domains/external).

## Setup

### From binaries

Download the most suitable binary from [the releases tab](https://github.com/aerialls/scaleway-ddns/releases). Builds are available for Linux, macOS and Windows.

```bash
tar xf scaleway-ddns_X.X.X_linux_amd64.tar.gz
./scaleway-ddns_X.X.X_linux_amd64/scaleway-ddns
```

### Using the docker image

```bash
docker run --rm -d  --name scaleway-ddns -v `pwd`:/config aerialls/scaleway-ddns:latest --config /config/scaleway-ddns.yml
```

## Usage

```bash
./scaleway-ddns --config /etc/scaleway-ddns/scaleway-ddns.yml --dry-run --verbose
```

## Parameters

* `--config` - Specify the location of the configuration file (**required**)
* `--dry-run` - Do not perform update actions (default `false`)
* `--verbose` - Display debug messages (default `false`)
* `--help` - Display the help message block

## Configuration

```yaml
scaleway:
  organization_id: __ORGANIZATION_ID__
  access_key: __ACCESS_KEY__
  secret_key: __SECRET_KEY__

domain:
  name: contoso.com
  record: public
  ttl: 60

interval: 300

ipv4:
  enabled: true
  url: https://api-ipv4.ip.sb/ip

ipv6:
  enabled: true
  url: https://api-ipv6.ip.sb/ip
```

**Note**: IPv6 is disabled by default.
