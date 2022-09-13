package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)


func ExampleScrape(n int,w *csv.Writer) {
	// Request the HTML page.
	sprintf := fmt.Sprintf("%d", n)
	res, err := http.Get("http://www.hahait.com/beijing/search.html?key2=%E5%8C%97%E4%BA%AC&page="+sprintf)
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

	// Find the review items
	doc.Find("div.clist li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find(".vipname").Text()
		if band == "" {
			band = s.Find(".name").Text()
		}

		cinfo := s.Find(".cinfo").Text()
		info := assInfo(cinfo)

		info.Company = band
		info.WriteExcel(w)
		//fmt.Printf("Review %d: 公司- %s 联系人- %s 地址- %s 经营范围- %s\n", i, info.Company, info.Contacts,info.Address,info.ScopeOfBusiness)
	})
}

func main() {
	f, err := os.Create("哈哈IT-北京.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	w.Write([]string{"公司名称", "联系人", "地址", "经营范围"})

	for i := 1; i <= 272; i++ {
		fmt.Printf("第%d页\n",i)
		ExampleScrape(i,w)
	}

	fmt.Println("end")

}



func assInfo(info string)  *CInfo{
	return &CInfo{
		Contacts:        contactsRegexp(info),
		Address:         address(info),
		ScopeOfBusiness: scopeOfBusiness(info),
	}
}

type CInfo struct {
	Company string
	Contacts string
	Address string
	ScopeOfBusiness string
}

func (c *CInfo) WriteExcel(w *csv.Writer) {
	w.Write([]string{c.Company,c.Contacts,c.Address,compressStr(c.ScopeOfBusiness)})
	w.Flush()
}



// 手机号正则
func iphoneRegexp(iphone string) []string{
	reg1 := regexp.MustCompile(`\d{11}`)
	return  reg1.FindAllString(iphone, -1)
}

//联系人
func contactsRegexp(cInfo string) string{
	contacts := ""
	s := ""

	if strings.Contains(cInfo,"联系人") {
		s = strings.Split(cInfo, "联系人：")[1]
	}else {
		s = strings.Split(cInfo, "地址：")[0]
	}


	if strings.Contains(s,"查看QQ号") {
		contacts = strings.Split(s, "查看QQ号")[0]
	}else {
		contacts = strings.Split(s, "地址：")[0]
	}


   return contacts
}

//业务范围
func scopeOfBusiness(cInfo string) string  {
	return  strings.Split(cInfo, "业务范围：")[1]
}


//地址
func address(cInfo string) string {
	s := strings.Split(cInfo, "地址：")[1]
	return strings.Split(s, "业务范围：")[0]

}

func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

func pedanticReadAll(r io.Reader) (b []byte, err error) {
	var bufa [64]byte
	buf := bufa[:]
	for {
		n, err := r.Read(buf)
		if n == 0 && err == nil {
			return nil, fmt.Errorf("Read: n=0 with err=nil")
		}
		b = append(b, buf[:n]...)
		if err == io.EOF {
			n, err := r.Read(buf)
			if n != 0 || err != io.EOF {
				return nil, fmt.Errorf("Read: n=%d err=%#v after EOF", n, err)
			}
			return b, nil
		}
		if err != nil {
			return b, err
		}
	}
}