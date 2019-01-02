package main

import (
	"gamersky/pkg/getc"
	"gamersky/pkg/setting"
	"log"
)

func main() {
	var err error

	// 初始化配置
	setting.Setup()

	r := getc.Reuqester{
		Url: "https://www.gamersky.com/ent/wp/",
	}

	err = r.RequestPage()

	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Println(r.BaseData)

}
