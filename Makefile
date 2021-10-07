build:
	go build -o bin/scws-words .
start:	
	nohup /data/scws/bin/scws-words &
stop:
	killall scws-words
test:
	curl "http://127.0.0.1:8020/words?key=关于幼教体系组织结构调整等的通知"
	curl "http://127.0.0.1:8020/words?key=外国钱币硬币银铌世界纸钞爱藏"
	curl "http://127.0.0.1:8020/words?key=矮人火枪地狱兽残酷角斗士的军刺"
reload:
	curl "http://127.0.0.1:8020/reload"
