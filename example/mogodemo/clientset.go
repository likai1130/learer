package main

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	/*client, err := typed.NewForConfigAndClient(nil)
	if err != nil {
		panic(err)
	}*/
	get()
	var stuesM []bson.M
	m := bson.M{}
	m["name"] = "222"
	stuesM = append(stuesM, m)
	var stues []Stu
	err := bsonToObjs(stuesM, &stues)
	if err != nil {

		panic(err)
	}
	fmt.Println(stues)
}

type Counter interface {
	Count() int
}

type Stu struct {
	Name string
}
func (s *Stu) Count() int {
	return 5
}

type Teacher struct {
	Counter
}

func get()  {
	teacher := Teacher{
		&Stu{},
	}
	fmt.Println(teacher.Count())
}

func bsonToObjs(bm []bson.M, objs interface{}) error {
	marshal, err := json.Marshal(bm)
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal,objs)
}