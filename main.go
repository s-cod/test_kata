package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {

	text := "IV * X"
	text = strings.Replace(text, " ", "", -1)

	var currentOperation string
	var typeArgs string

	for _, i := range Operation {
		if strings.Contains(text, i) {
			if currentOperation == "" {
				currentOperation = i
			} else {
				log.Fatal("Не верны входные данные: " + text)
			}
		}
	}

	if currentOperation == "" {
		log.Fatal("Не верны входные данные: " + text)
	}

	args := strings.Split(text, currentOperation)

	if len(args) != 2 {
		log.Fatal("Не верны входные данные: " + text)
	}

	typeArgs = SetTypeArgs(args)

	if typeArgs == "rim" {
		args[0] = fmt.Sprintf("%v", slices.Index(Rim, args[0])+1)
		args[1] = fmt.Sprintf("%v", slices.Index(Rim, args[1])+1)
	}

	result := 0

	switch currentOperation {
	case "+":
		result = Atoi(args[0]) + Atoi(args[1])
	case "-":
		result = Atoi(args[0]) - Atoi(args[1])
	case "/":
		if Atoi(args[1]) == 0 {
			log.Fatal("На 0 делить нельзя: " + text)
		}
		result = Atoi(args[0]) / Atoi(args[1])
	case "*":
		result = Atoi(args[0]) * Atoi(args[1])
	default:

	}
	if typeArgs == "arabic" {
		fmt.Println(result)
	}

	if typeArgs == "rim" {

		rimResult := ""
		if result < 1 || result > 99 {
			log.Fatal("Неверное римское число")
		}
		if result < 10 {
			t, _ := strconv.Atoi(fmt.Sprint(result))
			fmt.Println(Rim[t-1])
			return
		}
		numbers := fmt.Sprintf("%v", result)

		t, _ := strconv.Atoi(string(numbers[0]))

		for i := 0; i < t; i++ {
			rimResult = rimResult + "X"
		}
		t, _ = strconv.Atoi(string(numbers[1]))
		if t != 0 {
			rimResult = rimResult + Rim[t-1]
		}

		fmt.Println(rimResult)
	}

}
