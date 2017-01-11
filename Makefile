DIR_CONF=/etc/godoremi
DIR_SRV=/etc/systemd/system
SRV_NAME=godoremi.service

build:
		go build

clean:
		sudo systemctl stop $(SRV_NAME)
		sudo systemctl disable $(SRV_NAME)
		sudo rm $(DIR_SRV)/$(SRV_NAME)
		sudo rm -rf $(DIR_CONF)

install:
		go build
		sudo mkdir -p $(DIR_CONF)
		sudo cp config.json $(DIR_CONF)
		sudo ./godoremi -install
		sudo systemctl start $(SRV_NAME)
		sudo systemctl status $(SRV_NAME)

reinstall: clean install

test:
		go test ./...
