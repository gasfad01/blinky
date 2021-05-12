package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

func printHeader() {
	var header string = `
	
______ __________       ______          
___  /____  /__(_)_________  /______  __
__  __ \_  /__  /__  __ \_  //_/_  / / /
_  /_/ /  / _  / _  / / /  ,<  _  /_/ / 
/_.___//_/  /_/  /_/ /_//_/|_| _\__, /  
                               /____/ by Gasfad   

	`
	color.Red(header)
}

var target string

func main() {
	//Specify target
	if len(os.Args) > 1 {
		target = os.Args[1]
	} else {
		fmt.Println("Please specify target as an argument -> Example: blinky https://domain.com")
		os.Exit(0)
	}

	//Fetching URL
	resp, err := http.Get(target)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	defer resp.Body.Close()

	printHeader()
	fmt.Println("URL: ", target)
	fmt.Println("Response Code: ", resp.Status)
	fmt.Println("")

	//Parsing Domain
	domain, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	//Scraping
	c := colly.NewCollector(colly.AllowedDomains(domain.Hostname()), colly.Async(true))

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("%s\n", link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit(target)
	c.Wait()
}
