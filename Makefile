check:
	gofmt -e . > /dev/null

test:
	go test -v ./...