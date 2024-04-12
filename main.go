package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите значения для расчета: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	calc(text)
}

func calc(s string) {

	text := s
	text = strings.Replace(text, " ", "", -1)

	var currentOperation string
	var typeArgs string

	for _, i := range Operation {
		if strings.Contains(text, i) {
			if currentOperation == "" {
				currentOperation = i
			} else {
				panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			}
		}
	}

	if currentOperation == "" {
		panic("Выдача паники, так как строка не является математической операцией.")
	}

	args := strings.Split(text, currentOperation)

	if len(args) != 2 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	typeArgs, err := SetTypeArgs(args)
	if err != nil {
		panic(err)
	}

	result := 0
	if typeArgs == RimType {
		args[0] = fmt.Sprintf("%v", slices.Index(Rim, args[0])+1)
		args[1] = fmt.Sprintf("%v", slices.Index(Rim, args[1])+1)
	}

	switch currentOperation {
	case "+":
		result = Atoi(args[0]) + Atoi(args[1])
	case "-":
		result = Atoi(args[0]) - Atoi(args[1])
	case "/":
		if Atoi(args[1]) == 0 {
			panic("На 0 делить нельзя: " + text)
		}
		result = Atoi(args[0]) / Atoi(args[1])
	case "*":
		result = Atoi(args[0]) * Atoi(args[1])
	default:
	}

	if typeArgs == ArabicType {
		fmt.Println(result)
	}

	if typeArgs == RimType {
		rimResult := ""
		if result < 1 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		if result < 10 {
			t, _ := strconv.Atoi(fmt.Sprint(result))
			fmt.Println(Rim[t-1])
			return
		}
		numbers := fmt.Sprintf("%v", result)

		t, _ := strconv.Atoi(string(numbers[0]))
		if result < 50 {
			for i := 0; i < t; i++ {
				rimResult = rimResult + "X"
			}
			t, _ = strconv.Atoi(string(numbers[1]))
			if t != 0 {
				rimResult = rimResult + Rim[t-1]
			}
			fmt.Println(rimResult)
			return
		}
		if result == 90 {
			fmt.Println("XC")
			return
		}
		if result == 100 {
			fmt.Println("C")
			return
		}
		t = t - 5
		rimResult = "L"
		for i := 0; i < t; i++ {
			rimResult = rimResult + "X"
		}
		t, _ = strconv.Atoi(string(numbers[1]))
		if t != 0 {
			rimResult = rimResult + Rim[t-1]
		}
		fmt.Println(rimResult)
		return
	}
}
