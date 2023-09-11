package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Define the structure for the RSS feed.
type RSS struct {
	Channel *Channel `xml:"channel"`
	XmlName xml.Name `xml:"rss"`
}

// Define the structure for the channel within the RSS feed.
type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

// Define the structure for an individual item in the RSS feed.
type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"ht:news_item"`
}

// Define the structure for a news item within an item.
type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS

	// Read the trends data from Google Trends RSS feed.
	data := readTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Iterate through the trending items and print information.
	for i, item := range r.Channel.ItemList {
		rank := i + 1
		fmt.Printf("%d. %s\n", rank, item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Traffic: %s\n\n", item.Traffic)
		fmt.Println("------------------------------")

	}
}

// getTrends sends an HTTP GET request to the Google Trends RSS feed and returns the response.
func getTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return resp
}

// readTrends reads the data from the HTTP response body and returns it as a byte slice.
func readTrends() []byte {
	resp := getTrends()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return data
}
