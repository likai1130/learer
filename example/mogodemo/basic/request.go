package basic

import "go.mongodb.org/mongo-driver/bson"

type Result struct {
	id string
	count int64
	data []byte
	err error
}

func (r Result) Into (obj interface{}) error {
	if r.err != nil {
		return r.err
	}
	if err := bson.Unmarshal(r.data, obj); err != nil {
		return err
	}
	return nil
}

func (r Result) Count () int64 {
	return r.count
}

func (r Result) Id (obj interface{}) string {
	return r.id
}

func (r Result) Error (obj interface{}) error {
	return r.err
}