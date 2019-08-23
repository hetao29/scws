scws:
	cd scws-c && 
build:
	export GOPATH=`pwd`"/go/" && go build -o bin/test .
start:	
	./bin/test
