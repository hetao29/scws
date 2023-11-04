package scws

import (
	"goscws"
	"log"
	"os"
)

func init() {
}
func New(conf Configuration) *Info {
	info := new(Info)
	info.config = conf
	info.scws.New()
	return info
}

type Info struct {
	scws   goscws.Scws
	config Configuration
}

func file_is_exists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
func (this *Info) Set() {
	for _, v := range this.config.Dict {
		if file_is_exists(v.File) {
			if v.Type == "XDB" {
				this.scws.AddDict(v.File, goscws.SCWS_XDICT_XDB|goscws.SCWS_XDICT_MEM)
			} else if v.Type == "TXT" {
				this.scws.AddDict(v.File, goscws.SCWS_XDICT_TXT|goscws.SCWS_XDICT_MEM)
			} else {
				log.Fatal("error: dict type is not supports", v.Type)
			}
			log.Println("notice: add dict file: ", v.Type, v.File)
		} else {
			log.Fatal("can't open dict file: ", v.Type, v.File)
		}
	}
	if file_is_exists(this.config.Rule) {
		log.Println("notice: set rule file: ", this.config.Rule)
		this.scws.SetRule("/data/scws/dict/rules.utf8.ini")
	} else {
		log.Println("warning: not set rule file / or rule file not exists", this.config.Rule)
	}
	this.scws.SetCharset("utf8")
	this.scws.SetIgnore(1)
	this.scws.SetMulti(goscws.SCWS_MULTI_SHORT)
}
func (this *Info) Get(key string) []string {
	scws_fork, _ := this.scws.Fork()
	scws_fork.SendText(key)
	//var words []string;
	words := make([]string, 0, 100)
	for scws_fork.Next() {
		//fmt.Println(scws_fork.GetRes())
		words = append(words, scws_fork.GetRes().String)
	}
	scws_fork.Free()
	return words
}
func (this *Info) Free() {
	this.scws.Free()
}
