# goprops
A Go library to read and write Java properties file.

### How to Get It
```
go get github.com/fredyw/goprops
```

### Usage
```go
package main

import (
	"fmt"
	"os"
	"github.com/fredyw/goprops"
)

func testRead() {
	props, err := goprops.ReadFile("in.properties")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	for key, value := range props {
		fmt.Println(key, "=", value)
	}
}

func testWrite() {
	props := goprops.Properties{}
	props["key1"] = "value1"
	props["key2"] = "value2"
	props["key3"] = "value3"
	err := goprops.WriteFile("out.properties", props)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	testRead()
	testWrite()
}
```
