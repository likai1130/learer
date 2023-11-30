package tokens

import (
	"bufio"
	"fmt"
	"github.com/pkoukk/tiktoken-go"
	"io"
	"log"
	"os"
	"testing"
)

const encoding = "cl100k_base"

// main
func TestTokenNumbers(t *testing.T) {
	/*	list := ReadTestFile()
		testTokenByEncoding(list)*/
	//getTokenByEncoding(list, model)

	file, err := os.Open("likai3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(b))
	tokens := getTokenByEncoding(string(b), encoding)
	fmt.Println(tokens)
}

// read all columns from a file
func ReadTestFile() (textList []string) {
	file, err := os.Open("test/demo/likai.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	textList = lines
	return
}

// getTokenByEncoding
func getTokenByEncoding(text string, encoding string) (num_tokens int) {
	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf(": %v", err)
		return
	}
	token := tke.Encode(text, nil, nil)
	return len(token)
}

// 循环计算
// testTokenByEncoding
func testTokenByEncoding(textList []string) {
	for i := 0; i < len(textList); i++ {
		fmt.Printf("text: %s, encoding: %s, token: %d \n", textList[i], encoding, getTokenByEncoding(textList[i], encoding))
	}
}

// 一次性计算
