default:
		go build
install:
		go build
		sudo cp -Rf ./files/* /
		sudo ./godoremi -install
test:
		go test ./...