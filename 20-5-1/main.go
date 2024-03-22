//Напишите функцию, вычисляющую определитель матрицы размером 3 × 3.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func create(rows int, cols int) [][]int {
	m := make([][]int, rows)
	for k := 0; k < len(m); k++ {
		m[k] = make([]int, cols)
	}
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m[i][j] = rand.Intn(10)
		}
	}
	return m
}

func det(m [][]int) int {
	result := m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) - m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) + m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])
	return result
}

func main() {
	matrix := create(3, 3)
	fmt.Println("Сгенерированная матрица:")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("Определитель:", det(matrix))
}
