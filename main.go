package main

import (
	"fmt"
	"strconv"
	"strings"
)

func RomanToInt(s string) int {
	var RomanMap = map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	result := 0
	for R := range s {
		if R < len(s)-1 && RomanMap[s[R:R+1]] < RomanMap[s[R+1:R+2]] {
			result -= RomanMap[s[R:R+1]]
		} else {
			result += RomanMap[s[R:R+1]]
		}
	}
	return result
}

func IntegerToRoman(I int) string {
	conversions := []struct {
		key   int
		value string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{8, "VIII"},
		{7, "VII"},
		{6, "VI"},
		{5, "V"},
		{4, "IV"},
		{3, "III"},
		{2, "II"},
		{1, "I"},
	}
	var roman strings.Builder
	for _, conversion := range conversions {
		for I >= conversion.key {
			roman.WriteString(conversion.value)
			I -= conversion.key
		}
	}
	return roman.String()
}

const first = "Значения не могут быть больше 10"
const second = "Значения не могут быть меньше 0"
const third = "Значения не могут быть равны 0"
const fourth = "Неправильный ввод операции"
const fifth = "Результат римского выражения не может быть меньше 0"
const sixth = "Результат не может быть меньше 1"
const seventh = "Результат римского выражения не может быть равен 0"
const eighth = "Значения должны быть одного типа"
const ninth = "не правильный вид значений"

func main() {
	fmt.Println("Введите выражение")
	var x, op, y string
	fmt.Scan(&x, &op, &y)
	x1, _ := strconv.Atoi(x)
	y1, _ := strconv.Atoi(y)
	rome1 := RomanToInt(x)
	rome2 := RomanToInt(y)
	res1 := x1 + y1
	if res1 > 0 && x1 != 0 && y1 != 0 {
		result1 := x1 + y1
		result2 := x1 - y1
		result3 := x1 * y1
		result4 := x1 / y1
		if x1 > 10 || y1 > 10 {
			panic(first)
		} else if x1 < 0 || y1 < 0 {
			panic(second)
		} else if x1 == 0 || y1 == 0 {
			panic(third)
		} else if op != "+" && op != "-" && op != "*" && op != "/" {
			panic(fourth)
		} else if x1 > 0 && y1 > 0 && x1 <= 10 && y1 <= 10 && op == "+" {
			fmt.Println(result1)
		} else if x1 > 0 && y1 > 0 && x1 <= 10 && y1 <= 10 && op == "-" {
			fmt.Println(result2)
		} else if x1 > 0 && y1 > 0 && x1 <= 10 && y1 <= 10 && op == "*" {
			fmt.Println(result3)
		} else if x1 > 0 && y1 > 0 && x1 <= 10 && y1 <= 10 && op == "/" {
			fmt.Println(result4)
		}
	} else if res1 == 0 && rome1 > 0 && rome2 > 0 {
		resultRome1 := rome1 + rome2
		resultRome2 := rome1 - rome2
		resultRome3 := rome1 * rome2
		resultRome4 := rome1 / rome2
		int1 := IntegerToRoman(resultRome1)
		int2 := IntegerToRoman(resultRome2)
		int3 := IntegerToRoman(resultRome3)
		int4 := IntegerToRoman(resultRome4)
		if rome1 > 10 || rome2 > 10 {
			panic(first)
		} else if op != "+" && op != "-" && op != "*" && op != "/" {
			panic(fourth)
		} else if rome1 < rome2 && op == "-" {
			panic(fifth)
		} else if rome1 < rome2 && op == "/" {
			panic(sixth)
		} else if rome1 == rome2 && op == "-" && rome1 != 0 && rome2 != 0 {
			panic(seventh)
		} else if rome1 > 0 && rome2 > 0 && rome1 <= 10 && rome2 <= 10 && op == "+" {
			fmt.Println(int1)
		} else if rome1 > 0 && rome2 > 0 && rome1 <= 10 && rome2 <= 10 && op == "-" {
			fmt.Println(int2)
		} else if rome1 > 0 && rome2 > 0 && rome1 <= 10 && rome2 <= 10 && op == "*" {
			fmt.Println(int3)
		} else if rome1 > 0 && rome2 > 0 && rome1 <= 10 && rome2 <= 10 && op == "/" {
			fmt.Println(int4)
		}
	} else if res1 < 0 || x1 < 0 || y1 < 0 {
		panic(second)
	} else if res1 == 0 && (x1 < 0 || y1 < 0) && op == "+" {
		panic(second)
	} else if (res1 == x1 || res1 == y1) && (x1 != 0 || y1 != 0) {
		panic(eighth)
	} else if res1 == 0 && rome1 == 0 && rome2 == 0 {
		panic(ninth)
	}
}
