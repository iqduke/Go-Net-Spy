package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/opesun/goquery"
	"io/ioutil"
	"net/http"
	"os"
)

type webOperation interface {
	filter()
	getCharpterList()
}

type webUrl struct {
	name   string
	uri    string
	method string
	domain string
	//charpter []string
}

func (wl *webUrl) filter() {

}

func main() {

	wl := &webUrl{
		name:   "择天记",
		uri:    "http://www.xs82.com/books/31/31565/?LMCL=V6pzEW",
		method: "GET",
		domain: "http://www.xs82.com/books/31/31565/",
		//charpter: make([]string),
	}

	wl.getCharpterList()
}

func (wl *webUrl) getCharpterList() {
	client := &http.Client{}
	request, err := http.NewRequest(wl.method, wl.uri, nil)
	// request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	// request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	// request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	// request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	// request.Header.Set("Cache-Control", "max-age=0")
	// request.Header.Set("Connection", "keep-alive")
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(0)
	}

	response, err := client.Do(request)

	defer response.Body.Close()

	body, err1 := ioutil.ReadAll(response.Body)

	if err1 != nil {
		fmt.Println("body err")
	}

	by := string(body)
	fmt.Println(by)
	x, _ := goquery.ParseString(by)

	cp := x.Find("ul.chapterlist li a")
	fmt.Println("------------中文------------------")

	for i := 0; i < cp.Length(); i++ {
		d := cp.Eq(i).Attr("href")
		v := cp.Eq(i).Text()
		fmt.Println("href=" + d + " ----- value=" + v)
	}

}
