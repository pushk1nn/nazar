export CGO_ENABLED=1

build: 
	go build -o nazar main.go && \
	sudo setcap cap_net_raw,cap_net_admin=eip ./nazar