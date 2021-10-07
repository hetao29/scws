package scws

import (
	_"fmt"
	"os"
	"goscws"
)

func init() {
}
func New() *Info {
	info := new(Info)
	info.scws.New()
	return info
}

type Info struct {
	scws goscws.Scws
}

func file_is_exists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
func (this *Info) Set() {
	this.scws.AddDict("/data/scws/dict/dict.utf8.xdb", goscws.SCWS_XDICT_XDB | goscws.SCWS_XDICT_MEM)
	var file="/data/scws/dict/user.utf8.ini";
	if(file_is_exists(file)){
		this.scws.AddDict(file, goscws.SCWS_XDICT_TXT | goscws.SCWS_XDICT_MEM)
	}
	this.scws.SetRule("/data/scws/dict/rules.utf8.ini")
	this.scws.SetCharset("utf8")
	this.scws.SetIgnore(1)
	//this.scws.SetMulti(goscws.SCWS_MULTI_SHORT & goscws.SCWS_MULTI_DUALITY & goscws.SCWS_MULTI_ZMAIN)
	//this.scws.SetMulti(goscws.SCWS_XDICT_MEM | goscws.SCWS_MULTI_SHORT | goscws.SCWS_MULTI_DUALITY)
	this.scws.SetMulti(goscws.SCWS_MULTI_SHORT)
}
func (this *Info) Reload() {
	this.scws.SetDict("/data/scws/dict/dict.utf8.xdb", goscws.SCWS_XDICT_XDB | goscws.SCWS_XDICT_MEM)
	var file="/data/scws/dict/user.utf8.ini";
	if(file_is_exists(file)){
		this.scws.AddDict(file, goscws.SCWS_XDICT_TXT | goscws.SCWS_XDICT_MEM)
	}
	this.scws.SetRule("/data/scws/dict/rules.utf8.ini")
	this.scws.SetCharset("utf8")
	this.scws.SetIgnore(1)
	//this.scws.SetMulti(goscws.SCWS_MULTI_SHORT & goscws.SCWS_MULTI_DUALITY & goscws.SCWS_MULTI_ZMAIN)
	//this.scws.SetMulti(goscws.SCWS_XDICT_MEM | goscws.SCWS_MULTI_SHORT | goscws.SCWS_MULTI_DUALITY)
	this.scws.SetMulti(goscws.SCWS_MULTI_SHORT)
}
func (this *Info) Get(key string) []string {
	scws_fork, _ := this.scws.Fork()
	scws_fork.SendText(key)
	//var words []string;
	words :=make([]string,0,100);
	for scws_fork.Next() {
		//fmt.Println(scws_fork.GetRes())
		words=append(words,scws_fork.GetRes().String)
	}
	return words;
}
func (this *Info) Free() {
	this.scws.Free()
}
