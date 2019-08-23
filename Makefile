build:
	export GOPATH="/data/scws/go/" && go build -o bin/test .
start:	
	./bin/test
