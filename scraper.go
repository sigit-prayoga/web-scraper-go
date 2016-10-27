package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//get arguments
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Add a URL to proceed!")
		return
	}
	scrap(args[0])
}

func scrap(url string) {
	fmt.Printf("Scraping... : %s \n", url)
	//get the HTML DOM
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	//get a detail container
	contentBox := doc.Find("div#prodinfo")
	//get all the details like title, desc and price
	detail := getDetail(contentBox)
	fmt.Println(detail)
}

func getDetail(el *goquery.Selection) ProductDetail {
	//get title
	titleBox := el.Find("h1#prod_title")
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
	discount := el.Find("span#product_saving_percentage").Text()
	afterDiscount := el.Find("span#special_price_box").Text()
	//construct Price object
	price = Price{normal: normal, discount: discount, afterDiscount: afterDiscount}
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
