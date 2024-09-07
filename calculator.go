package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToIntMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRomanMap = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
}

func romanToInt(s string) (int, error) {
	if val, ok := romanToIntMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("некорректное римское число")
}

func intToRoman(num int) (string, error) {
	if num > 0 && num <= 10 {
		return intToRomanMap[num], nil
	}
	return "", fmt.Errorf("некорректный результат для римской системы")
}

func calculate(input string) (string, error) {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	var operator string

	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	} else {
		return "", fmt.Errorf("некорректная операция")
	}

	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		return "", fmt.Errorf("некорректный формат операции")
	}

	isRoman := false
	a, err1 := strconv.Atoi(parts[0])
	b, err2 := strconv.Atoi(parts[1])

	if err1 != nil && err2 != nil {
		isRoman = true
		a, err1 = romanToInt(parts[0])
		b, err2 = romanToInt(parts[1])
		if err1 != nil || err2 != nil {
			return "", fmt.Errorf("некорректные римские числа")
		}
	} else if err1 != nil || err2 != nil {
		return "", fmt.Errorf("используются одновременно разные системы счисления")
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", fmt.Errorf("деление на ноль")
		}
		result = a / b
	}

	if isRoman {
		if result <= 0 {
			return "", fmt.Errorf("в римской системе нет отрицательных чисел или нуля")
		}
		romanResult, err := intToRoman(result)
		if err != nil {
			return "", err
		}
		return romanResult, nil
	}

	return strconv.Itoa(result), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите математическую операцию (например, 1 + 2 или VI / III):")
	for scanner.Scan() {
		input := scanner.Text()
		result, err := calculate(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
		fmt.Println("Введите следующую операцию:")
	}
}
