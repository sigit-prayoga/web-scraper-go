# Web Scraper with GO
Scraping web page with [GO](https://golang.org). I'm scraping [Lazada](http://www.lazada.co.id/) as one of the largest ecommerce site in Indonesia.


# Usage
Go to [Lazada](http://www.lazada.co.id/) and find any product and copy the link. 

*For example:*

[http://www.lazada.co.id/polo-carion-330006-trifungsi-hitam-8594153.html](http://www.lazada.co.id/polo-carion-330006-trifungsi-hitam-8594153.html)

I'm assuming you have already installed `go` in your machine and set `$GOPATH`. Otherwise, please refer to this [installation guide](https://golang.org/doc/install).

## 1. Install `goquery`
```sh
$ go get github.com/PuerkitoBio/goquery
```

## 2. Navigate to project folder
```sh
$ cd path/to/app
```

## 3. Build the app
```sh
$ go build scraper.go
```

## 4. Run the app with URL that you copied
```sh
$ ./scraper yourURL
```

## 5. Simpler way to do it (Skip #2 - #4)
```sh
$ go run scraper.go yourURL
```