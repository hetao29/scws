# 说明

基于[scws](https://github.com/hightman/scws/)的golang的开源的中文分词系统，特点：
1. 支持编译安装，纯golang代码，无需要第3方的静态和动态库（和scws已经混合编码）
2. docker / docker swarm / k8s 部署
3. 支持多个自定义词库，格式可以和txt也可以是xdb格式（参考 [scws](https://github.com/hightman/scws/) 的词典生成工具scws-gen-dict）
4. 支持规则文件
5. etc/dict下面可以把百度输入法的词库转换成txt分词词库的工具
6. rest api接口进行分词，也支持实时reload词库操作
具体可以参考etc下的配置文件，或者Makefile的示例

## Docker & Swarm/Composer
```bash
docker pull hetao29/scws:latest
```

## 编译
```bash
make build
```

## 运行

```bash
make start
```

## 自定义字典(格式)

1. <词>  <词频(TF)>  <词重(IDF)>  <词性(北大标注)>
2. 用空格或者制表符分开,tf,idf,attr不是必须
3. 词性参考 http://www.xunsearch.com/scws/docs.php#attr

```ini
#<word>[\t<tf>[\t<idf>[\t<attr>]]]
```

## 测试

```bash
make test
```
结果
```json
{"message":"pong","words":["关于","幼教","体系","组织","结构调整","结构","调整","等","的","通知"]}curl "http://127.0.0.1:8020/words?key=外国钱币硬币银铌世界纸钞爱藏"
{"message":"pong","words":["外国","钱币","硬币","银","铌","世界","纸钞","爱","藏"]}curl "http://127.0.0.1:8020/words?key=矮人火枪地狱兽残酷角斗士的军刺"
{"message":"pong","words":["矮人","火枪","地狱兽","地狱","残酷","角斗士","角斗","的","军","刺"]}
```
