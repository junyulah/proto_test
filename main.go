package main

import (
	"encoding/json"
	"fmt"

	"proto_test/types"

	"google.golang.org/protobuf/encoding/protojson"
)

type MyDetails struct {
	A int   `json:"a"`
	B []int `json:"b"`
}

func main() {
	payload := []byte(`{"message": "msg111", "details": {"a": 1, "b": [4,5,6]}}`)

	// 1. unmarshal payload to proto type
	var req types.Example
	err := protojson.Unmarshal(payload, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(req)

	// 2. convert proto.Value type to go struct

	// 2.1 get raw json
	bs, err := req.Details.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))

	// 2.2 raw json to go struct
	var myDetails MyDetails
	err = json.Unmarshal(bs, &myDetails)
	if err != nil {
		panic(err)
	}

	fmt.Println(myDetails)
}
