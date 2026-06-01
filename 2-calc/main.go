package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CalcFunc func([]float64) float64

var operations = map[string]CalcFunc{
	"AVG": calculateAvg,
	"SUM": calculateSum,
	"MED": calculateMed,
}

var operationNames = map[string]string{
	"AVG": "Среднее",
	"SUM": "Сумма",
	"MED": "Медиана",
}

func readValue() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	return strings.TrimSpace(str), err
}

func readOperation() string {
	for {
		fmt.Println("Выберите операцию: AVG (среднее), SUM (сумма), MED (медиана)")
		operation, err := readValue()
		if err != nil {
			fmt.Println("Ошибка чтения ввода.")
			continue
		}
		operation = strings.ToUpper(operation)

		if _, ok := operations[operation]; ok {
			return operation
		}
		fmt.Println("Неверная операция. Попробуйте снова.")
	}
}

func readNumbers() []float64 {
	for {
		fmt.Printf("Введите числа через запятую: ")
		str, err := readValue()
		if err != nil {
			fmt.Println("Ошибка чтения ввода.")
			continue
		}
		trimmedStr := strings.TrimSpace(str)
		if trimmedStr == "" {
			fmt.Println("Пустой ввод, попробуйте снова.")
			continue
		}

		parts := strings.Split(trimmedStr, ",")
		numbers := make([]float64, 0, len(parts))
		validInput := true

		for _, part := range parts {
			part = strings.TrimSpace(part)
			number, err := strconv.ParseFloat(part, 64)
			if err != nil {
				fmt.Printf("Ошибка: '%s' не является числом. Попробуйте снова.\n", part)
				validInput = false
				break
			}
			numbers = append(numbers, number)
		}

		if validInput {
			return numbers
		}
	}
}

func calculateAvg(numbers []float64) float64 {
	return calculateSum(numbers) / float64(len(numbers))
}

func calculateSum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calculateMed(numbers []float64) float64 {
	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Float64s(sortedNumbers)
	n := len(sortedNumbers)
	if n%2 == 1 {
		return sortedNumbers[n/2]
	}
	return (sortedNumbers[n/2-1] + sortedNumbers[n/2]) / 2
}

func main() {
	fmt.Println("Калькулятор операций над числами")

	for {
		opKey := readOperation()
		numbers := readNumbers()

		calcFunc := operations[opKey]
		result := calcFunc(numbers)
		opName := operationNames[opKey]

		fmt.Printf("%s: %.2f\n", opName, result)

		fmt.Print("Хотите выполнить еще одну операцию? (y/n): ")
		answer, _ := readValue()
		if strings.ToLower(answer) != "y" {
			break
		}
	}
}
