package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var err error

	imgSrc := "http://img1.gamersky.com/image2018/12/20181229_ddw_459_1/gamersky_01origin_01_201812291176E5.jpg"
	res, err := http.Get(imgSrc)

	if err != nil {
		fmt.Println(err.Error())
	}

	//doc, err := goquery.NewDocumentFromReader(res.Body)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	content, err := ioutil.ReadAll(res.Body)
	fmt.Printf("%v", content)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = ioutil.WriteFile("img.jpg", content, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

}
