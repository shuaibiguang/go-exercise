package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan *Message)
}

type LogProcess struct {
	read Reader  // 读取器
	write Writer // 写入器
	rc chan []byte
	wc chan *Message
}

type ReadFormFile struct {
	path string
}

type WriteToInfluxDB struct {
	influxDBDsn string
}

type Message struct {
	TimeLocal time.Time
	BytesSent int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime float64
}

//读取模块
func (this *ReadFormFile) Read (rc chan []byte) {
	// 首先打开文件
	// 从文件末尾开始读取
	// 将读取的内容传入通道当中
	if f, err := os.Open(this.path); err == nil {
		f.Seek(0, 2)
		rd := bufio.NewReader(f)
		for {
			if line, err := rd.ReadBytes('\n'); err == nil {
				//fmt.Println(string(line))
				rc <- line[:len(line)-1]
			} else if err == io.EOF {
				time.Sleep(1 * time.Second)
			} else if err != nil {
				panic(fmt.Sprintf("read file error: %s", err.Error()))
			}
		}
	} else {
		panic(fmt.Sprintf("open file error: %s", err.Error()))
	}
}

// 写入模块
func (this *WriteToInfluxDB) Write (wc chan *Message) {
	for v := range wc {
		fmt.Println(v)
	}
}

// 解析模块
func (this *LogProcess) Process () {
	/**
	172.0.0.12 - - [04/Mar/2018:13:49:52 +0000] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.854
	*/
	r := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)

	loc, _ :=time.LoadLocation("Asia/Shanghai")

	for v := range this.rc {
		ret := r.FindStringSubmatch(string(v))
		//log.Println(ret, len(ret))
		if len(ret) != 14 {
			log.Println("FindStringSubmatch fail:", string(v))
			continue
		}
		message := &Message{}

		/*--------------*/
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)

		if err != nil {
			log.Println("ParseInLocation fail:", err.Error(), ret[4])
		}
		/*-------------*/
		byteSent, _ := strconv.Atoi(ret[8])
		/*-------------*/
		// GET /foo?query=t HTTP/1.0
		reqSli := strings.Split(ret[6], " ")

		if len(reqSli) != 3 {
			log.Println("strings.Split fail :" , ret[6])
			continue
		}
		/*-------------*/
		u, err := url.Parse(reqSli[1])
		if (err != nil) {
			log.Println("url.Parse fail :", err.Error(), reqSli[1])
		}
		/*-------------*/

		upstreamTime,_ := strconv.ParseFloat(ret[12], 64)
		requestTime,_ := strconv.ParseFloat(ret[13], 64)

		message.TimeLocal = t
		message.BytesSent = byteSent
		message.Method = reqSli[0]
		message.Path = u.Path
		message.Scheme = ret[5]
		message.Status = ret[7]
		message.UpstreamTime = upstreamTime
		message.RequestTime = requestTime

		this.wc <- message
	}
}

func main() {
	//r := &ReadFormFile{path: "./access.log"}
	//w := &WriteToInfluxDB{influxDBDsn: "asdasdasdas"}
	lp := &LogProcess{
		read: &ReadFormFile{path: "./access.log"},
		write: &WriteToInfluxDB{influxDBDsn: "asdasdasdas"},
		rc: make(chan []byte),
		wc: make(chan *Message),
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(60 * time.Second)
}
