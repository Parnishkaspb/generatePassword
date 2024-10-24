package main

import (
	"Project1/internal/pkg/generatepassword"
	"flag"
	"fmt"
)

func main() {
	var includeUpperCase bool
	var includeNumbers bool
	var includeSpecialChars bool
	var length int
	var countRep int

	flag.BoolVar(&includeUpperCase, "up", false, "Использовать заглавные буквы")
	flag.BoolVar(&includeNumbers, "numbers", false, "Использовать цифры в пароле")
	flag.BoolVar(&includeSpecialChars, "spec", false, "Использовать специальные символы")
	flag.IntVar(&length, "length", 5, "Длина пароля")
	flag.IntVar(&countRep, "repeats", 1, "Количество паролей за раз")

	flag.Parse()

	generator := generatepassword.New()
	passwords := generator.GeneratePassword(includeUpperCase, includeNumbers, includeSpecialChars, length, countRep)

	fmt.Println("Сгенерированные пароли:")
	fmt.Println(passwords)
}
