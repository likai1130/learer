package main

import "fmt"

func foo(m map[string]int) {
	m["key1"] = 11
	m["key2"] = 12
}

/*func main() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	fmt.Println(m) // map[key1:1 key2:2]
	foo(m)
	fmt.Println(m) // map[key1:11 key2:12]
}*/
func main()  {
	var m =  map[string]int{}

	/*key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)

	fmt.Printf("The length of nil map: %d\n",
		len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n",
		key)
	delete(m, key)
	**/

	elem := 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem // 这里会引发panic。

}
