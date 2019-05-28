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

delete:
	rm -f ./constructor.yaml ./constructor.tpl
	rm -f ./structure/constructor.go

reset: delete setup
