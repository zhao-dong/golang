package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
{"timestamp":"2020-05-19T08:04:54.743+0000","status":403,"message":"Access Denied"}
*/

type SitemapIndex struct {
	Status       int
	Message      string
	Timestamp111 string `json:"timestamp"`
}

func (m SitemapIndex) String() string {
	return "The record is :" + string(m.Status) + m.Message + m.Timestamp111
}

func main() {
	resp, _ := http.Get("https://jljzcar.com.cn/auth/login")
	//resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	//string_body := string(bytes)
	//fmt.Println(string_body)

	var s SitemapIndex
	json.Unmarshal(bytes, &s)
	fmt.Println(s)
	fmt.Println(s.Timestamp111)
}
