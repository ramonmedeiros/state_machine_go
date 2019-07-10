default: tests


tests:
	go test -v ./...

format-code:
	go fmt 
