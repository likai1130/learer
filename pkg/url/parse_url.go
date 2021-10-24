package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func CheckWebsite(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return false
	}

	if response.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func main() {
	//我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
	s := "https://pnode.solarfs.io/dn/short/b4e321144937ca55e85e79b76241beac-程序员5.jpg/"
	matched, err := regexp.Match("(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]", []byte(s))
	fmt.Println(matched,err)


	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}


	//直接访问 scheme。
	fmt.Println(u.Scheme)
	//User 包含了所有的认证信息，这里调用 Username和 Password 来获取独立值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	//Host 同时包括主机名和端口信息，如过端口存在的话，使用 strings.Split() 从 Host 中手动提取端口。
	fmt.Println(u.Host)
	if strings.Contains(u.Host, ":") {
		h := strings.Split(u.Host, ":")
		fmt.Println(h[0])
		fmt.Println(h[1])
	}

	//这里我们提出路径和查询片段信息。
	fmt.Println(u.Path)

	fmt.Println(string(u.RequestURI()))
	if u.Path != "" {
		s = strings.Split(s, u.Path)[0]
		fmt.Println(s)
	}

	i := s[len(s)-1:]
	fmt.Println(i)
	website := CheckWebsite(s)
	fmt.Println(website)

}
