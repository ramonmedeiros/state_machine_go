default: tests


tests:
	go test -v ./...

format-code:
	go fmt github.com/ramonmedeiros/state_machine_go/states
	go fmt github.com/ramonmedeiros/state_machine_go/users
