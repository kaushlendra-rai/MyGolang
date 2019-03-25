package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAgg struct {
	Title string
	News  map[string]NewsMap
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Root")
}

func aggregateHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	data := NewsAgg{"My First News title", getNewsMap()}
	tmpl := template.Must(template.ParseFiles("layouts/newsAgg.html"))
	tmpl.Execute(w, data)

	fmt.Println("Total time to fetch aggregate news::", time.Now().Sub(start))
}
func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/agg", aggregateHandler)
	http.ListenAndServe(":80", nil)
}

func getNewsFromLocation(c chan News, Location string) {
	defer wg.Done()
	var n News
	// Somehow the Location has a trailing '\n' in the string which causes GET to fail. Hence we need to trim it.
	Location = strings.Trim(Location, "\n")
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &n)

	c <- n
}

//parallel fetch : 21.5657073s(30 buffer channels), 20.0463272s(100 buffer channels), 20.1646985s(500 channels)
func getNewsMap() map[string]NewsMap {
	var siteMap SitemapIndex

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &siteMap)
	resp.Body.Close()

	queue := make(chan News, 20)
	news_map := make(map[string]NewsMap)
	for _, Location := range siteMap.Locations {
		wg.Add(1)
		go getNewsFromLocation(queue, Location)
	}
	wg.Wait()
	close(queue)

	for n := range queue {
		for idx, title := range n.Titles {
			news_map[title] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	return news_map
}

// Linear fetch time: 1m52.9508763s, 1m57.0500733
func getNewsMap_Linear() map[string]NewsMap {
	var siteMap SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &siteMap)
	resp.Body.Close()

	news_map := make(map[string]NewsMap)
	for _, Location := range siteMap.Locations {
		// Somehow the Location has a trailing '\n' in the string which causes GET to fail. Hence we need to trim it.
		Location = strings.Trim(Location, "\n")
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		xml.Unmarshal(bytes, &n)

		for idx, title := range n.Titles {
			news_map[title] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	return news_map
}
