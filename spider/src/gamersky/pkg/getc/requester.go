package getc

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

type Reuqester struct {
	Url string // 请求的url,
	BaseData string
	PageFormatData string
	Query interface{}
}

func (this *Reuqester) Init () error {
	res, err := http.Get(this.Url)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Status code not 200, is :%d", res.StatusCode)
	}

	//使用query格式化页面数据
	if doc, err := goquery.NewDocumentFromReader(res.Body); err != nil {
		return err
	} else {
		this.Query = doc
	}

	// 逐行读取页面基础数据
	this.BaseData = ""

	rd := bufio.NewReader(res.Body)
	for {
		if line, _, err := rd.ReadLine(); err == nil {
			this.BaseData = this.BaseData + string(line)
		} else if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Read page data error: %v", err)
			return err
		}
	}

	return nil
}

//func (this *Reuqester) RequestPage () error {
//	res, err := http.Get(this.Url)
//
//	if err != nil {
//		return err
//	}
//
//	if res.StatusCode != 200 {
//		return fmt.Errorf("Status code not 200, is :%d", res.StatusCode)
//	}
//
//	rd := bufio.NewReader(res.Body)
//
//	// 逐行读取页面内容
//	this.BaseData = ""
//	for {
//		if line, _, err := rd.ReadLine(); err == nil {
//			this.BaseData = this.BaseData + string(line)
//		} else if err == io.EOF {
//			break
//		} else if err != nil {
//			log.Println("Read page data error: %v", err)
//			return err
//		}
//	}
//
//	return nil
//
//}

