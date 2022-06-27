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

	output := strings.Join(phrases[:], "; ")
	r, g, b := 255, 215, 0 // gold color

	for j := 0; j < len(output); j++ {
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
}
