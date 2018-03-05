package main

import (
	"fmt"
	"github.com/Nhoya/go-bing"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//Create new Client instance
	client := bing.NewClient("INSERT TOKEN HERE")
	//Start searching
	resp, err := client.Search("Nhoya gOSINT")
	if err != nil {
		panic(err)
	}
	//Iterate over search results
	for _, result := range resp.WebPages.Value {
		fmt.Println(result.Name, "||", result.URL)
	}
}
