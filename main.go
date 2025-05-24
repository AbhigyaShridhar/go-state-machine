package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	_ = errors.New("something went wrong")
}
