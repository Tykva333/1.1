package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Читаем матрицу из файла
	matrix, err := readMatrix("input.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Проверяем что матрица квадратная
	if !isSquare(matrix) {
		fmt.Println("Ошибка: матрица должна быть квадратной")
		return
	}

	// Вычисляем нужные значения
	det := determinant(matrix)
	trace := matrixTrace(matrix)
	transposed := transpose(matrix)

	// Сохраняем результаты
	err = saveResults("output.txt", det, trace, transposed)
	if err != nil {
		fmt.Println("Ошибка при сохранении:", err)
		return
	}

	// Выводим результаты
	printResults(det, trace, transposed)
}

func readMatrix(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		row := make([]float64, len(parts))

		for i, part := range parts {
			num, err := strconv.ParseFloat(part, 64)
			if err != nil {
				return nil, err
			}
			row[i] = num
		}

		matrix = append(matrix, row)
	}

	return matrix, scanner.Err()
}

func isSquare(matrix [][]float64) bool {
	size := len(matrix)
	for _, row := range matrix {
		if len(row) != size {
			return false
		}
	}
	return true
}

func determinant(matrix [][]float64) float64 {
	size := len(matrix)

	if size == 1 {
		return matrix[0][0]
	}

	if size == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	det := 0.0
	sign := 1.0

	for col := 0; col < size; col++ {
		minor := getMinor(matrix, 0, col)
		det += sign * matrix[0][col] * determinant(minor)
		sign *= -1
	}

	return det
}

func getMinor(matrix [][]float64, rowToRemove, colToRemove int) [][]float64 {
	size := len(matrix)
	minor := make([][]float64, size-1)

	i := 0
	for row := 0; row < size; row++ {
		if row == rowToRemove {
			continue
		}

		minor[i] = make([]float64, size-1)
		j := 0

		for col := 0; col < size; col++ {
			if col == colToRemove {
				continue
			}
			minor[i][j] = matrix[row][col]
			j++
		}
		i++
	}

	return minor
}

func matrixTrace(matrix [][]float64) float64 {
	trace := 0.0
	for i := 0; i < len(matrix); i++ {
		trace += matrix[i][i]
	}
	return trace
}

func transpose(matrix [][]float64) [][]float64 {
	size := len(matrix)
	result := make([][]float64, size)

	for i := range result {
		result[i] = make([]float64, size)
		for j := range result[i] {
			result[i][j] = matrix[j][i]
		}
	}

	return result
}

func saveResults(filename string, det, trace float64, matrix [][]float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintf(writer, "Определитель: %.2f\n", det)
	fmt.Fprintf(writer, "След: %.2f\n", trace)
	fmt.Fprintln(writer, "Транспонированная матрица:")

	for _, row := range matrix {
		for _, val := range row {
			fmt.Fprintf(writer, "%.2f ", val)
		}
		fmt.Fprintln(writer)
	}

	return nil
}

func printResults(det, trace float64, matrix [][]float64) {
	fmt.Println("Результаты:")
	fmt.Printf("Определитель: %.2f\n", det)
	fmt.Printf("След: %.2f\n", trace)
	fmt.Println("Транспонированная матрица:")
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%.2f ", val)
		}
		fmt.Println()
	}
}