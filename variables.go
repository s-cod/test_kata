package main

import (
	"log"
	"slices"
	"strconv"
)

var Rim = []string{
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"VIII",
	"IX",
	"X",
}

var Arabic = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
}

var Operation = []string{
	"+",
	"-",
	"/",
	"*",
}

func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Неверныее вычисляемые значения")
	}
	return result
}

func SetTypeArgs(args []string) string {

	if slices.Contains(Rim, args[0]) {
		if slices.Contains(Rim, args[1]) {
			return "rim"
		} else {
			log.Fatal("Не верны входные данные: ")
		}

	}
	if slices.Contains(Arabic, args[0]) {
		if slices.Contains(Arabic, args[1]) {
			return "arabic"
		} else {
			log.Fatal("Не верны входные данные: ")
		}
	}

	return ""
}
