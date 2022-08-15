package main

import "fmt"

// go tool compile -S -N -l main.go
type Art struct {
	Name string
	Size int
	Labels map[string]interface{}
}



func main()  {
	m := make(map[string]interface{})
	m["xx"] = 111

	var art = &Art{Name: "kli", Size: 30,Labels: m}
	fmt.Printf("init art=%+v\n",art)
	Update(art)
	fmt.Printf("updated art=%+v\n",art)
}

func Asset(a Art) {
	s := a


	s.Labels["4444"] = 10
	s.Size = 20

}
func Update(a *Art)  {
	Asset(*a)
	a.Name = "李四"
}