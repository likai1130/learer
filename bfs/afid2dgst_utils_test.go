package dgstseed

import (
	"log"
	"testing"
)

func TestGenAfid(t *testing.T) {
	afid, err := GetAfidLocal("")
	if err != nil {
		panic(err)
	}

	dgst := ConvertAfig2Dgst(afid)
	log.Println(dgst)
}
