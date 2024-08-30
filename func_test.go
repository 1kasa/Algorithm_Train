package algorithmtrain

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type a struct {
	Function interface{} `json:"Function"`
}

type Printfer interface {
	Printf()
}

type b struct {
	b  int
	bs string
}

func (bb *b) Printf() {
	log.Printf("b:%d,bs:%s\n\n", bb.b, bb.bs)
}

type c struct {
	c  int
	cs string
}

func (cc *c) Printf() {
	log.Printf("c:%d,cs:%s\n\n", cc.c, cc.cs)
}

type d struct {
	d  int
	ds string
}

func (dd *d) Printf() {
	log.Printf("d:%d,ds:%s\n\n", dd.d, dd.ds)
}

func TestFunction(t *testing.T) {
	jsonStr := `{"Function":{"b":10, "bs":"efwgfewfw"}}`

	var AAAA a
	err := json.Unmarshal([]byte(jsonStr), &AAAA)
	if err != nil {
		panic(err)
	}
    
	if funcMap, ok := AAAA.Function.(map[string]interface{}); ok {
		if _, exists := funcMap["b"]; exists {
			var bObj b
			bBytes, _ := json.Marshal(funcMap)
			_ = json.Unmarshal(bBytes, &bObj)
			bObj.Printf()
		} else if _, exists := funcMap["c"]; exists {
			var cObj c
			cBytes, _ := json.Marshal(funcMap)
			_ = json.Unmarshal(cBytes, &cObj)
			cObj.Printf()
		} else if _, exists := funcMap["d"]; exists {
			var dObj d
			dBytes, _ := json.Marshal(funcMap)
			_ = json.Unmarshal(dBytes, &dObj)
			dObj.Printf()
		} else {
			fmt.Println("Unknown type")
		}
	} else {
		fmt.Println("No map[string]interface{}")
	}
}
