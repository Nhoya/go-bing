# go-bing
Golang Library for bing  search API v7

Before getting started with this lib you need to get a [Bing Search Token](https://azure.microsoft.com/it-it/try/cognitive-services/my-api)

`go-bin` is written with simplicity in mind

# Installation
`go get -v github.com/Nhoya/go-bing`

# Usage

```
package main

import (
        "fmt"
        "github.com/Nhoya/go-bing"
)

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
            //Printing result name and URL    
            fmt.Println(result.Name, "||", result.URL)
        }
}

//Output:
//GitHub - Nhoya/gOSINT: OSINT framework in Go || https://github.com/Nhoya/gOSINT
//Refactoring lots of stuff · Nhoya/gOSINT@2bb5d91 · GitHub || https://github.com/Nhoya/gOSINT/commit/2bb5d913da7c3ad4636cfc02c719e347a6c2fc2e
//gOSINT - Open Source Intelligence Framework - KitPloit ... || https://www.kitploit.com/2018/01/gosint-open-source-intelligence.html
//gOSINT – Open Source Threat Intelligence Gathering ... || https://pentesttoolz.com/2018/01/22/gosint-open-source-threat-intelligence-gathering-processing-framework-2/
//...
```

