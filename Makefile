build:
	export GOPATH=`pwd`"/go/" CGO_LDFLAGS="-I/data/scws/scws-lib/include/ -L/data/scws/scws-lib/lib/ -lscws" && go build -o bin/scws-words .
buildscws:
	git submodule init && git submodule update
	cd scws-c && aclocal && libtoolize && automake && autoconf && autoheader
	cd scws-c && ./configure --prefix=`pwd`/"../scws-lib" && make && make install
start:	
	export LD_LIBRARY_PATH=/data/scws/scws-lib/lib/ && nohup /data/scws/bin/scws-words &
stop:
	killall scws-words
