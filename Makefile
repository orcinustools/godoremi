default:
		go build
install:
		go build
		sudo mkdir -p /etc/godoremi
		sudo cp config.json /etc/godoremi
		sudo ./godoremi -install
test:
		go test ./...