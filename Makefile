default: tests


test:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...

format-code:
	go fmt github.com/ramonmedeiros/state_machine_go/states
	go fmt github.com/ramonmedeiros/state_machine_go/users
