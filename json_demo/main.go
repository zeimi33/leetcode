package main

import (
	"encoding/json"
	"net/url"
)
import "fmt"
import "encoding/base64"
type BaseInfo struct {
	Machinenumber string `json:"machinenumber"`
	Machinename string `json:"machinename"`
	Machineip string `json:"machineip"`
	Machinetype string `json:"machinetype"`
}
func main() {
	params := &BaseInfo{
		Machinenumber: "159ABBE56E057F20F883E",
		Machinename:   "yangbin-pc",
		Machineip:     "192.168.0.122",
		Machinetype:   "Skpzkpfrq",
	}
	a := "JTdCJTIybWFjaGluZW51bWJlciUyMiUzQSUyMjE1OUFCQkU1NkUwNTdGMjBGODgzRSUyMiUyQyUyMm1hY2hpbmVuYW1lJTIyJTNBJTIyeWFuZ2Jpbi1wYyUyMiUyQyUyMm1hY2hpbmVpcCUyMiUzQSUyMjE5Mi4xNjguMC4xMjIlMjIlMkMlMjJtYWNoaW5ldHlwZSUyMiUzQSUyMlNrcHprcGZycSUyMiU3RA=="
	decodeBytes, err := base64.StdEncoding.DecodeString(a)
	fmt.Println(string(decodeBytes))
	jsonStr, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	b := url.PathEscape(string(jsonStr))
	b64Str := base64.StdEncoding.EncodeToString([]byte(b))
	fmt.Println(b64Str)
}