package getc

import (
	"fmt"
	"gamersky/pkg/setting"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

type Reuqester struct {
	Url            string // 请求的url,
	PageFormatData string
	Body           []io.ReadCloser
	Query          goquery.Document
}

func (this *Reuqester) Init() error {
	// 使用chan 进行阻塞
	<-setting.Interval

	log.Printf("正在请求: %s", this.Url)

	res, err := http.Get(this.Url)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Status code not 200, is :%d，url：%s", res.StatusCode, this.Url)
	}

	log.Println("请求成功，正在格式化网页数据")

	defer res.Body.Close()

	//使用query格式化页面数据
	if doc, err := goquery.NewDocumentFromReader(res.Body); err != nil {
		return err
	} else {
		this.Query = *doc
	}
	return nil
}
