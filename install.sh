#!/bin/bash

# Install Go
wget https://go.dev/dl/go1.25.3.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.25.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Dependencies for packet decoding
sudo apt install libpcap-dev gcc
export CGO_ENABLED=1

# Configure env variables
export IFACE=$1
echo "[+] Set to listen on interface $1"

# Build and set binary capabilities
go build -o nazar main.go && sudo setcap cap_net_raw,cap_net_admin=eip ./nazar