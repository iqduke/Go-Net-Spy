package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iqduke/Go-Net-Spy/sample"
	"github.com/opesun/goquery"
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

func getContent(url string, method string) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, nil)
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
}

func main() {

	sample.Run()

	wl := &webUrl{
		name:   "择天记",
		uri:    "http://www.xs82.com/books/31/31565/?LMCL=GmCrrf",
		method: "GET",
		domain: "http://www.xs82.com/books/31/31565/",
		//charpter: make([]string),
	}

	//wl.getCharpterList()

	getContent(wl.domain+"11885686.html", wl.method)
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
	//fmt.Println(by)
	x, _ := goquery.ParseString(by)

	cp := x.Find("ul.chapterlist li a")
	fmt.Println("------------中文------------------")

	db, dberr := sql.Open("mysql", "root:1qaz@WSX@/noval")

	if dberr != nil {
		panic(dberr.Error())
	}

	defer db.Close()

	stmtIns, stmterr := db.Prepare("insert into bookdetail (urlhtml,content,createtime) values(?,?,?)")

	if stmterr != nil {
		panic(stmterr.Error())
	}
	defer stmtIns.Close()

	for i := 0; i < cp.Length(); i++ {
		d := cp.Eq(i).Attr("href")
		//v := cp.Eq(i).Text()

		_, err = stmtIns.Exec(d, "", time.Now()) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

	}

	fmt.Println("end")
}
