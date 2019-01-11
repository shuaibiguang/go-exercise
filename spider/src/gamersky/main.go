package main

import (
	"gamersky/pkg/getc"
	"gamersky/pkg/setting"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
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

	if err = r.Init(); err != nil {
		log.Println(err.Error())
		return
	}

	// 内容列表页，拿取到所有2级网页的url，并且进行下一页访问
	go func() {
		r.Query.Find(".contentpaging li").Each(func(i int, selection *goquery.Selection) {
			listUrl, b := selection.Find(".con a").Attr("href")
			if b {
				setting.ListUrl <- listUrl
			}
		})
	}()

	// 进入图片列表页，拿取图片url， 并且进行下一页访问，拿取所有的url
	go func() {
		for url := range setting.ListUrl {
			r2 := getc.Reuqester{
				Url: url,
			}
			if err := r2.Init(); err != nil {
				log.Println(err.Error())
				continue
			}
			r2.Query.Find(".Mid2L_con p[align='center'] a").Each(func(i int, selection *goquery.Selection) {
				imgurl, b := selection.Attr("href")
				if b {
					// 查询到图片地址后，进行切割拿取图片Url
					kv := strings.Split(imgurl, "?")
					if len(kv) == 2 {
						setting.ImgSrc <- kv[1]
					}
				}
			})

			// 查找是否有下一月， 如果有的话，存入chan 继续追踪下一页
			r2.Query.Find(".page_css a").Each(func(i int, selection *goquery.Selection) {
				if selection.Text() == "下一页" {
					nextUrl, _ := selection.Attr("href")
					setting.ListUrl <- nextUrl
				}
			})

		}
	}()

	// 下载图片
	go func() {
		for url := range setting.ImgSrc {
			res, err := http.Get(url)

			if err != nil {
				log.Println(err.Error())
				continue
			}

			if res.StatusCode != 200 {
				log.Printf("Status code not 200, is :%d，url: %s", res.StatusCode, url)
				continue
			}

			content, err := ioutil.ReadAll(res.Body)

			if err != nil {
				log.Println(err.Error())
				continue
			}
			// 获取文件名
			fileNames := strings.Split(url, "/")
			fileName := fileNames[len(fileNames)-1]

			log.Printf("正在下载文件，文件名：%s", fileName)

			err = ioutil.WriteFile(setting.CSetting.DownloadPath+fileName, content, 0777)

			if err != nil {
				log.Println(err.Error())
			}
		}
	}()

	// 限制 请求间隔
	go func() {
		timers := 0
		for {
			if timers >= setting.CSetting.IntervalTime {
				timers = 0
				setting.Interval <- 2
			} else {
				timers++
			}
			time.Sleep(time.Second * 1)
		}
	}()

	block := make(chan int)
	<-block
}
