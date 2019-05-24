PROJECT?=constructor
PACKAGE?=github.com/$(PROJECT)

install: 
	go install

build:
	go build

test: 
	go test $(PACKAGE)... 

help: 
	constructor --help

container:
	echo "TODO"

