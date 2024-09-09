package algorithmtrain

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type a struct {
	Function Printfer `json:"Function"`
}

type Printfer interface {
	Printf()
}

type b struct {
	B  int    `json:"b"`
	Bs string `json:"bs"`
}

func (bb *b) Printf() {
	log.Printf("b: %d, bs: %s\n", bb.B, bb.Bs)
}

type c struct {
	C  int    `json:"c"`
	Cs string `json:"cs"`
}

func (cc *c) Printf() {
	log.Printf("c: %d, cs: %s\n", cc.C, cc.Cs)
}

type d struct {
	D  int    `json:"d"`
	Ds string `json:"ds"`
}

func (dd *d) Printf() {
	log.Printf("d: %d, ds: %s\n", dd.D, dd.Ds)
}

func(afunc *a)UnmarshalJson(data []byte)error {
    var tmp map[string]interface{}

	if err := json.Unmarshal(data, &tmp); err!= nil {
		return err
	}
	if functionData, ok := tmp["Function"]; ok {
       var bbb b
	   functionJSON, err := json.Marshal(functionData)
	   if err != nil {
		   return err
	   }
	   if err := json.Unmarshal(functionJSON,&bbb);err == nil {
		 afunc.Function = &bbb
		 return nil
	   }else {
		   var ccc c
		   functionJSON, err := json.Marshal(functionData)
		   if err != nil {
			   return err
		   }
		   if err := json.Unmarshal(functionJSON,&ccc);err == nil {
			   afunc.Function = &ccc
			   return nil
		   }else {
			   var ddd d
			   functionJSON, err := json.Marshal(functionData)
			   if err != nil {
				   return err
			   }
			   if err := json.Unmarshal(functionJSON,&ddd);err == nil {
				   afunc.Function = &ddd
				   return nil
			   }else {
				   fmt.Println("error,Unknow function")
			   }
		   }
	   }
	}
	return fmt.Errorf("error Function")
}

func TestFunction(t *testing.T) {
	jsonStr := `{"Function":{"b":10, "bs":"efwgfewfw"}}`

	var aaa a
	err := aaa.UnmarshalJson([]byte(jsonStr))
	if err != nil {
		t.Fatal("UnmarshalJson failed")
	}
    
	aaa.Function.Printf()
}