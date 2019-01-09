package main

import (
	"gamersky/pkg/getc"
	"gamersky/pkg/setting"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"
)


func main() {
	var err error

	// 初始化配置
	setting.Setup()

	setting.Interval <- 1

	r := getc.Reuqester{
		Url: "https://www.gamersky.com/ent/wp/",
	}

	err = r.Init()

	if err != nil {
		log.Printf("%v", err)
		return
	}

	// 内容列表页，拿取到所有2级网页的url，并且进行下一页访问
	go func () {
		r.Query.Find(".contentpaging li").Each(func(i int, selection *goquery.Selection) {
			listUrl, b := selection.Find(".con a").Attr("href")
			if b {
				setting.ListUrl<-listUrl
			}
		})
	} ()

	// 进入图片列表页，拿取图片url， 并且进行下一页访问，拿取所有的url
	go func () {
		for url := range setting.ListUrl {
			r2 := getc.Reuqester{
				Url: url,
			}
			r2.Init()
			r2.Query.Find(".Mid2L_con p[align='center'] a").Each(func(i int, selection *goquery.Selection) {
				imgurl, b := selection.Attr("href")
				if b {
					// 查询到图片地址后，进行切割拿取图片Url
					kv := strings.Split(imgurl, "?")
					if len(kv) == 2 {
						log.Println(kv[1])
						setting.ImgSrc <- kv[1]
					}
				}
			})
		}
	} ()

	// 限制 请求间隔
	go func () {
		timers := 0
		for {
			log.Println(timers, setting.CSetting.IntervalTime)
			if timers >= setting.CSetting.IntervalTime {
				timers = 0
				setting.Interval <- 2
			} else {
				timers++
			}
			time.Sleep(time.Second * 1)
		}
	} ()

	time.Sleep(time.Second * 30)
}
