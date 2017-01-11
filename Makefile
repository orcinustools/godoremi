build:
		go build

clean:
		sudo systemctl stop godoremi.service
		sudo systemctl disable godoremi.service
		sudo rm -rf /etc/godoremi
		sudo rm /etc/systemd/system/godoremi.service

install:
		go build
		sudo mkdir -p /etc/godoremi
		sudo cp config.json /etc/godoremi
		sudo ./godoremi -install
		sudo systemctl start godoremi.service
		sudo systemctl status godoremi.service

test:
		go test ./...
