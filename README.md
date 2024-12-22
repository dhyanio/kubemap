[![release](https://img.shields.io/github/release/dhyanio/kubemap.svg)](https://github.com/dhyanio/kubemap/releases)
[![tests](https://github.com/dhyanio/kubemap/actions/workflows/test.yaml/badge.svg)](https://github.com/dhyanio/kubemap/actions/workflows/test.yaml)
[![linter](https://github.com/dhyanio/kubemap/actions/workflows/linter.yaml/badge.svg)](https://github.com/dhyanio/kubemap/actions/workflows/linter.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dhyanio/kubemap)](https://goreportcard.com/report/github.com/dhyanio/kubemap)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.23-61CFDD.svg?style=flat-square)
[![release](https://godoc.org/github.com/dhyanio/kubemap?status.svg)](https://pkg.go.dev/github.com/dhyanio/kubemap?tab=doc)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

Kubemap is a tfstate visualizer.

# Table of Contents
- [Installation](#installation)
- [Cli](#cli)

### Installation

#### Kubemap

**From releases**
This installs binary.

* Linux
```
curl -LO "https://github.com/dhyanio/kubemap/releases/download/$(curl -s https://api.github.com/repos/dhyanio/kubemap/releases/latest | grep tag_name | cut -d '"' -f 4)/kubemap-linux-amd64"
chmod +x kubemap-linux-amd64
sudo mv kubemap-linux-amd64 /usr/local/bin/kubemap
```
* MacOS
```
curl -LO "https://github.com/dhyanio/kubemap/releases/download/$(curl -s https://api.github.com/repos/dhyanio/kubemap/releases/latest | grep tag_name | cut -d '"' -f 4)/kubemap-darwin-amd64"
chmod +x kubemap-darwin-amd64
sudo mv kubemap-darwin-amd64 /usr/local/bin/kubemap
```

**From source**
1.  Run `git clone <kubemap repo> && cd kubemap/`
2.  Run `make build`
