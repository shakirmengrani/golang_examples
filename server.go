package main

import (
	"fmt"
	"net/http"
	"html/template"
	"io/ioutil"
	"encoding/xml"
	"sync"
)

var wg sync.WaitGroup
var news_map map[string]NewsMap = make(map[string]NewsMap)


type NewsAggPage struct{
	Title string
	News map[string]NewsMap
}


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

func getNewsRoutine(newChannel chan News,location string){
	defer wg.Done()
	var news News
	resp, _ := http.Get(location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &news)
	for idx, _ := range news.Titles{
		news_map[news.Titles[idx]] = NewsMap{Keyword: news.Keywords[idx], Location: news.Locations[idx]}
	}
	newChannel <- news
}


func newsAgg(w http.ResponseWriter, r * http.Request){
	var s SitemapIndex
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	var newsChannel = make(chan News, len(s.Locations))
	for _, Location := range s.Locations{
		wg.Add(1)
		go getNewsRoutine(newsChannel, Location)
	}
	wg.Wait()
	close(newsChannel)
	for elem := range newsChannel{
		for idx, _ := range elem.Titles{
			news_map[elem.Titles[idx]] = NewsMap{Keyword: elem.Keywords[idx], Location: elem.Locations[idx]}
		}
	}		
	p := NewsAggPage{"Title goes here", news_map}
	t, _ := template.ParseFiles("index.html")
	err := t.Execute(w, p)
	if err != nil{
		fmt.Println(err)
	}
}

func index_handler(w http.ResponseWriter, r * http.Request){
	fmt.Fprintf(w, "Hello, web from Go !")
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/news", newsAgg)
	http.ListenAndServe(":8000", nil)
}