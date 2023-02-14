//
// main.go
// Copyright (C) 2023 rmelo <rmelo@r-melo-lnx>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {

	encoded := strings.Join(os.Args[1:], " ")
	text, err := url.QueryUnescape(encoded)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(text)
}
