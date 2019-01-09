package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var ListUrl = make(chan string)
var ImgSrc = make(chan string, 100)
var Interval = make(chan int, 1)

type C struct {
	DownloadPath string
	IntervalTime int
}

var CSetting = &C{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/c.ini")

	if err != nil {
		log.Fatalf("Fail to parse 'conf/c.ini':%v", err)
	}

	err = cfg.MapTo(&CSetting)

	if err != nil {
		log.Fatalf("Cfg.MapTo ReadSetting err: %v", err)
	}
	//CSetting.IntervalTime = CSetting.IntervalTime * time.Second
}
