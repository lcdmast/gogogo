package main

import "fmt"

//interface
type VowelsFinds interface {
	FindVowels() []rune
}

type MyString string

//
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

func main() {
	name := MyString("lichangdong")
	var v VowelsFinds
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}
