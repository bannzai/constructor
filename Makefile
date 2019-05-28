PROJECT?=constructor
PACKAGE?=github.com/$(PROJECT)

install: 
	go install

test: 
	go test ./...

setup: 
	constructor setup

dry-run:
	constructor generate

reset: 
	rm -f ./constructor.yaml ./constructor.tpl
