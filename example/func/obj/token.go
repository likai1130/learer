package main

import (
	"fmt"
	"path/filepath"
	"strconv"
)

func main() {

	serialNo :=34464
	fmt.Printf("serialNo=%x \n",serialNo)

	hashPre:="xxxxxxxxxxx"
	tokenId:=hashPre + fmt.Sprintf("%04x",serialNo)

	fmt.Printf("tokenId=%s \n",tokenId)

	t := tokenId[0:len(tokenId) -4]
	fmt.Println(t)
	t2 := tokenId[len(tokenId) -4:len(tokenId)]
	fmt.Println(t2)

	parseInt, err := strconv.ParseInt(t2, 16, 64)
	if err != nil {
		panic(parseInt)
	}
	fmt.Println(parseInt)

	type student struct {
		name string
	}

	var student1 = &student{name: tokenId}
	student1 = nil
	fmt.Printf("%v",student1)

	var uploadDir = filepath.Join("configs.AppConfig.Server.DataPath", "upload")
	fmt.Println(uploadDir)


}
