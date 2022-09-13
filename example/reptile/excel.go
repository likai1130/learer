package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type person struct {
	JOB_NUMBER string
	NAME       string
	DEP_CODE   string
	DEP_NAME   string
	EMAIL      string
	IC_NUMBER  string
	EXTENSION  string
}

func CreateCaptcha() string {
	return fmt.Sprintf("%02v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100))
}

func main() {
	f, err := os.Create("test.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dep_code := []string{"A1000", "A1100", "A1110", "A1120", "A2000", "A2100", "A2110"}
	dep_name := []string{"人事", "财务", "总务", "劳安", "研发", "业务", "制造"}

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	w.Write([]string{"JOB_NUMBER", "NAME", "DEP_CODE", "DEP_NAME", "EMAIL", "IC_NUMBER", "EXTENSION"})

	for i := 1; i < 10000; i++ {
		var p person
		p.JOB_NUMBER = "1000" + strconv.Itoa(i)
		p.NAME = "test_" + strconv.Itoa(i)

		rand.Seed(time.Now().UnixNano())
		var dep_num int

		if i%3 == 0 {
			dep_num = 4
		} else {
			dep_num = rand.Intn(7)
		}
		p.DEP_CODE = dep_code[dep_num]
		p.DEP_NAME = dep_name[dep_num]
		p.EMAIL = p.NAME + "@test.com"

		p.IC_NUMBER = strconv.Itoa(i) + CreateCaptcha()
		p.EXTENSION = strconv.Itoa(i) + CreateCaptcha()

		w.Write([]string{p.JOB_NUMBER, p.NAME, p.DEP_CODE, p.DEP_NAME, p.EMAIL, p.IC_NUMBER, p.EXTENSION})
		w.Flush()

	}

}

