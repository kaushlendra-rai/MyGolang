package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"` // The ` used with xml is not a single quote but a symbol under the sign ~ on keyboard
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {

	resp, e := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	//data := string(bytes)
	// fmt.Println(data)
	if e != nil {
		fmt.Println(e)
	}

	var siteMap SitemapIndex
	xml.Unmarshal(bytes, &siteMap)

	for _, location := range siteMap.Locations {
		fmt.Printf("\n%s", location)
	}

	resp.Body.Close()
}
