package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
	"unsafe"
)

const (
	_ = iota
	Blue
	Red
	Yellow
)

func main()  {
	str := "李凯"
	
	fmt.Println(len(str))
	fmt.Printf("方法一原生rune类型：%d \n",len([]rune(str)))
	count := strings.Count(str, "") -1
	fmt.Printf("方法二 strings.count：%d \n",count)

	n := utf8.RuneCountInString(str)
	fmt.Printf("方法三utf8.RuneCountInString：%d \n",n)
	for k,v := range str{
		fmt.Printf("v type: %T index,val: %v,%v \n",v,k,v)
	}

	fmt.Println("------------------------------")
	for i:=0 ; i< len(str) ; i++{
		fmt.Printf("v type: %T index,val:%v,%v \n",str[i],i,str[i])
	}
}

/**

	一、 为什么len()方法输出不是字符串的字面长度？

		len()函数是用来获取字符串的字节长度，rune一个值代表的就是一个Unicode字符，所以求rune切片的长度就是字符个数。

	二、	为什么range关键字可以读取字符串的字面长度？

        使用下标遍历获取的是ASCII字符，而使用Range遍历获取的是Unicode字符。

	三、	Unicode和字符编码。

		1. 什么是Unicode？

			Unicode背景:

			Unicode把所有语言都统一到一套编码里。总结来说："unicode其实就是对字符的一种编码方式，可以理解为一个字符---数字的映射机制，利用一个数字即可表示一个字符。"

		2. 什么是字符编码？

			虽然unicode把所有语言统一到一套编码里了，但是他却没有规定字符对应的二进制码是如何存储。

			以汉字“汉”为例，它的 Unicode 码点是 0x6c49，对应的二进制数是 110110001001001，二进制数有 15位，这也就说明了它至少需要 2个字节来表示。可以想象，在Unicode 字典中往后的字符可能就需要 3个字节或者 4个字节，甚至更多字节来表示了。

			为了较好的解决Unicode 的编码问题， UTF-8 和UTF-16 两种当前比较流行的编码方式诞生了。UTF-8 是目前互联网上使用最广泛的一种Unicode编码方式，它的最大特点就是可变长。它可以使用 1 - 4个字节表示一个字符，根据字符的不同变换长度。在UTF-8编码中，一个英文为一个字节，一个中文为三个字节。

			UTF-8 是对Unicode的升级，一个英文为一个字节，一个中文为三个字节。

	四、Rune

			rune是int32的别名，在所有方面都等同于int32，按照约定，它用于区分字符值和整数值。

			说的通俗一点就是rune一个值代表的就是一个Unicode字符，因为一个Go语言中字符串编码为UTF-8，使用1-4字节就可以表示一个字符，所以使用int32类型范围就可以完美适配。


	五、总结

		1. Go语言源代码始终为UTF-8
		2. Go语言的字符串可以包含任意字节，字符底层是一个只读的byte数组。
		3. Go语言中字符串可以进行循环，使用下标循环获取的acsii字符，使用range循环获取的unicode字符。
		4. Go语言中提供了rune类型用来区分字符值和整数值，一个值代表的就是一个Unicode字符。
		5. Go语言中获取字符串的字节长度使用len()函数，获取字符串的字符个数使用utf8.RuneCountInString函数或者转换为rune切片求其长度，
			这两种方法都可以达到预期结果。

*/

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}

func main2() {
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将string类型变量地址显式转型为reflect.StringHeader
	fmt.Printf("0x%x\n", hdr.Data) // 0x10a30e0
	p := (*[5]byte)(unsafe.Pointer(hdr.Data)) // 获取Data字段所指向的数组的指针
	dumpBytesArray((*p)[:]) // [h e l l o ]   // 输出底层数组的内容
}