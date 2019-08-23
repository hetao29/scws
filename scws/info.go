package scws

import (
	_"fmt"
	"github.com/hetao29/goscws"
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

func (this *Info) Set() {
	this.scws.SetDict("/data/scws/dict/dict.utf8.xdb", goscws.SCWS_XDICT_XDB)
	this.scws.SetRule("/data/scws/dict/rules.utf8.ini")
	this.scws.SetCharset("utf8")
	this.scws.SetIgnore(1)
	//this.scws.SetMulti(goscws.SCWS_MULTI_SHORT & goscws.SCWS_MULTI_DUALITY & goscws.SCWS_MULTI_ZMAIN)
	//this.scws.SetMulti(goscws.SCWS_XDICT_MEM | goscws.SCWS_MULTI_SHORT | goscws.SCWS_MULTI_DUALITY)
	this.scws.SetMulti(goscws.SCWS_XDICT_MEM)
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
