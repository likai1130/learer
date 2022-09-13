package mutext

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

type S1 struct {
}

func (s1 S1) f() {
	fmt.Println("S1.f()")
}
func (s1 S1) g() {
	fmt.Println("S1.g()")
}

type S2 struct {
	S1
}

func (s2 S2) f() {
	fmt.Println("S2.f()")
}

type I interface {
	f()
}

func printType(i I) {
	if s1, ok := i.(S1); ok {
		s1.f()
		s1.g()
	}
	if s2, ok := i.(S2); ok {
		s2.f()
		s2.g()
	}
}

func TestPrint(t *testing.T) {
	printType(S1{})
	printType(S2{})
}

func TestGo(t *testing.T) {
	m := sync.Map{}
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			m.Store(rand.Int(),rand.Int())
		}()
	}
	wg.Wait()

}

type S struct {
	a, b, c string
}

func TestS(t *testing.T)  {
	//s := &S{"a", "b", "c"}
	x := interface{}(S{"a", "b", "c"})
	y := interface{}(S{"a", "b", "c"})
	fmt.Println(x == y)
}

type Result struct {
	Status int `json:"status"`
}

func TestR(t *testing.T) {
	var data = []byte(`{"status":200}`)
	result := &Result{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("result=%+v", result)
}


func TestD(t *testing.T) {
	fmt.Println(runtime.NumCPU())
	//fmt.Println(doubleScore(0))
	//fmt.Println(doubleScore(20.0))
	fmt.Println(doubleScore(50.0))
}
func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	score = source * 2
	return score
}