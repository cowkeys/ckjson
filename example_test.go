package ckjson_test

import (
	"github.com/cowkeys/ckjson"
)

func Example() {
	json := `{"employees": [
    { "firstName":"Bill" , "salary":10000.00,"flag":true },
    { "first-Name":"George" , "lastName":"Bush","flag":false },
    { "firstName":"Thomas" , "depart_tag":"IT" }]}`

	ck := ckjson.NewCkj("TestCK", json)
	ck.JsonToStruct()
}
