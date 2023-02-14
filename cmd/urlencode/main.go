//
// main.go
// Copyright (C) 2023 rmelo <rmelo@r-melo-lnx>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	text := strings.Join(os.Args[1:], " ")
	encoded := url.QueryEscape(text)
	fmt.Println(encoded)
}
