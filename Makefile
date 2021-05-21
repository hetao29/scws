build:
	export GOPATH=`pwd`"/go/" CGO_LDFLAGS="-I/data/scws/scws-lib/include/ -L/data/scws/scws-lib/lib/ -lscws" && go build -o bin/scws-words .
buildscws:
	git submodule init && git submodule update
	cd scws-c && aclocal && libtoolize && automake && autoconf && autoheader
	cd scws-c && autoreconf -vif && ./configure --prefix=`pwd`/"../scws-lib" && make && make install
start:	
	export LD_LIBRARY_PATH=/data/scws/scws-lib/lib/ && nohup /data/scws/bin/scws-words &
stop:
	killall scws-words
test:
	curl "http://127.0.0.1:8020/words?key=关于幼教体系组织结构调整等的通知"
	curl "http://127.0.0.1:8020/words?key=外国钱币硬币银铌世界纸钞爱藏"
	curl "http://127.0.0.1:8020/words?key=矮人火枪地狱兽残酷角斗士的军刺"
reload:
	curl "http://127.0.0.1:8020/reload"
