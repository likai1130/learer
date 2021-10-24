package main
import (
	"fmt"
	"sort"
)

type Person1 struct {
	Name string
	Age int
}


type PersonSlice []Person1

func (p PersonSlice) Len() int {
	return len(p)
}

func (p PersonSlice) Less(i, j int) bool {
	return p[i].Age > p[j].Age
}

func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PersonSlice) Sort() {
	sort.Sort(p)
}

func main() {
	instance := PersonSlice{{ "Jack",12}, {"Lucy",15}, { "Lilei",13}}
	instance.Sort()
	fmt.Println(instance) // [{12 Jack male} {13 Lilei male} {15 Lucy famale}]
}