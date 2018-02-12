package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type SitemapIndex struct{
	Locations []string `xml:"sitemap>loc"`
}

type News struct{
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

// type Location struct{
// 	Loc string `xml:"loc"`
// }

// func (l Location) String() string{
// 	return fmt.Sprintf(l.Loc)
// }

func main(){
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	for _, Location := range s.Locations{
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles{
			news_map[n.Titles[idx]] = NewsMap{Keyword: n.Keywords[idx], Location: n.Locations[idx]}
		}
	}
	for idx, _ := range news_map{
		fmt.Println(idx)
	}

}



