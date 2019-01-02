package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type C struct {
	PageFormat   string
	DownloadPath string
	IntervalTime time.Duration
}

var CSetting = &C{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/c.ini")

	if err != nil {
		log.Fatalf("Fail to parse 'conf/c.ini':%v", err)
	}

	err = cfg.MapTo(CSetting)

	if err != nil {
		log.Fatalf("Cfg.MapTo ReadSetting err: %v", err)
	}

	CSetting.IntervalTime = CSetting.IntervalTime * time.Second
}
