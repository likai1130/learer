package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	aa := "448f85064a238a1f53e54701463daa076ea208af948de0b63657e038525e59c363f69ae83f864a37b926f767a24a8bfba380f075e362806856e0ca5bd2470a12714258fd6bf49871f4959f5afd7e47bd7a4447617415691df7aa5d68c51390269b09830718fe13eea3a56ed15a9c3d7effd7fa71edaefb2ea46c56b39cbdb07a"
	bytes, _ := hex.DecodeString(aa)
	fmt.Println(string(bytes))

}