package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//1 from lazada, 0 from blibli
var from int

func main() {
	start := time.Now()
	//get arguments
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Add a URL to proceed!")
		return
	}
	scrap(args[0])
	total := time.Since(start)

	fmt.Printf("Time spent: %f secs \n", total.Seconds())
}

func scrap(url string) {
	fmt.Printf("Scraping... : %s \n", url)
	//get the HTML DOM
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	var productInfoSelector string
	//TODO change using enum
	if strings.Contains(url, "lazada.co.id") {
		//from lazada
		productInfoSelector = "div#prodinfo"
		from = 1
	} else if strings.Contains(url, "blibli.com") {
		//from blibli
		productInfoSelector = "section.product-detail-info"
		from = 0
	} else {
		// I dont care
		fmt.Println("Invalid URL, try using from Lazada.co.id or blibli.com")
		return
	}

	//get a detail container
	contentBox := doc.Find(productInfoSelector)
	//get all the details like title, desc and price
	detail := getDetail(contentBox)
	fmt.Println(detail)
}

func getDetail(el *goquery.Selection) ProductDetail {
	//get title
	var sel = "h1#prod_title"
	if from == 0 {
		sel = "h1"
	}
	titleBox := el.ChildrenFiltered(sel)
	title := titleBox.Text()

	//get price
	price := getPrice(el)

	//get description
	description := getDescription(el)

	//construct new product detail
	prodDetail := ProductDetail{title: title, description: description, price: price}

	return prodDetail
}

func getDescription(el *goquery.Selection) string {
	var description string
	//parse DOM
	el.Find("ul.prd-attributesList.ui-listBulleted.js-short-description").Each(func(index int, item *goquery.Selection) {
		item.Find("li").Each(func(i int, item *goquery.Selection) {
			desc := item.Find("span").Text()
			description += desc + " | "
		})
	})
	return description
}

func getPrice(el *goquery.Selection) Price {
	var price Price
	//parse DOM
	normal := el.Find("span#price_box").Text()
	if from == 0 {
		el.Find("span#strikeThroughPrice").Each(func(i int, el *goquery.Selection) {
		})
	}
	discount := el.Find("span#product_saving_percentage").Text()
	if from == 0 {
		discount = el.Find("span.price-discount").ChildrenFiltered("b").Text()
	}
	afterDiscount := el.Find("span#special_price_box").Text()
	if from == 0 {
		afterDiscount = el.Find("h2#priceDisplay").Text()
	}
	//construct Price object
	price = Price{normal: normal, discount: discount, afterDiscount: afterDiscount}
	fmt.Println(price)
	return price
}

type Price struct {
	normal        string
	discount      string
	afterDiscount string
}

type ProductDetail struct {
	title       string
	description string
	price       Price
}
