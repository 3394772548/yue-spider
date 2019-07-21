build:
	go build -tags "doc"
build_prod:
	go build
build_all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o yue-linux .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o yue-win.exe .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o yue-mac .