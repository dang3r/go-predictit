package main

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	body, err := ioutil.ReadFile("sample.json")
	if err != nil {
		log.Fatalf("Error reading file : %s", err)
	}

	var stuff Result
	err = json.Unmarshal(body, &stuff)
	if err != nil {
		log.Fatalf("Unable to parse body : %s", err)
	}

	data, _ := json.MarshalIndent(stuff, "", "  ")
	os.Stdout.Write(data)
	//	fmt.Printf("%v", data)
	//	fmt.Printf("%#v", stuff)
}
