package main

import (
	"fmt"
	"strings"

	"syreclabs.com/go/faker"
)

func main() {
	var phrases []string

	for i := 1; i < 3; i++ {
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	fmt.Println(strings.Join(phrases[:], "; "))
}
