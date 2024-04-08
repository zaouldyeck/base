package main

import (
	"fmt"

	"github.com/zaouldyeck/external"
)

func main() {
	s1 := external.Bark()
	s2 := external.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
}
