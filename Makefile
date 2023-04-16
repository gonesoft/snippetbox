.DEFAULT_GOAL := run

build: clear_bin
	GOARCH=amd64 GOOS=darwin go build -o bin/snippetbox_app github.com/gonesoft/snippetbox/cmd/web
	go test -v ./...
	./bin/snippetbox_app

run:
	go run github.com/gonesoft/snippetbox/cmd/web

clear_bin:
	go clean -cache
	rm -rf bin