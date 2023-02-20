package helper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func ExampleScrape(n int) {
	// Request the HTML page.
	url := fmt.Sprintf("http://tjgpc.zwfwb.tj.gov.cn/basSupplierInfo/getSupplierInfoList.do?page=%d&pagesize=10", n)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("div.cur").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		s.Find("a").Remove()
		str := s.Find("td").Text()
		if strings.Contains(str, "[供应商列表]") || strings.Contains(str, "审核通过,") {
			str = strings.ReplaceAll(str, "审核通过,", "")
		}

		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, "\t", "")
		str = strings.ReplaceAll(str, "司", "司\n")
		str = strings.TrimSpace(str)
		fmt.Println(str)
	})
}
