#!/bin/bash

apt-get update
apt-get install -y git bash bash-completion less nano zsh curl make jq ca-certificates build-essential

curl -sS https://starship.rs/install.sh | sh -s -- -y
echo 'eval "$(starship init bash)"' >> ~/.bashrc

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.3.1

echo "Setup completed!"