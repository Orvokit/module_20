//Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.

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
			m[i][j] = rand.Intn(5)
		}
	}
	return m
}

func multiplication(m1 [][]int, m2 [][]int) [3][4]int {

	var result [3][4]int
	//result[0][0] = m1[0][0] * m2[0][0] + m1[0][1] * m2[1][0] + m1[0][2] * m2[2][0]
	//result[0][1] = m1[0][0] * m2[0][1] + m1[0][1] * m2[1][1] + m1[0][2] * m2[2][1]
	//result[0][2] = m1[0][0] * m2[0][2] + m1[0][1] * m2[1][2] + m1[0][2] * m2[2][2]
	//result[ ][ ] =
	for k := 0; k < len(m1); k++ {
		for i := 0; i < len(m2[0]); i++ {
			for j := 0; j < len(m2); j++ {
				result[k][i] += m1[k][j] * m2[j][i]
			}
		}
	}
	return result
}

func main() {
	matrix1 := create(3, 5)
	matrix2 := create(5, 4)

	fmt.Println("Сгенерированная матрица 1:")
	for i := 0; i < len(matrix1); i++ {
		fmt.Println(matrix1[i])
	}

	fmt.Println("Сгенерированная матрица 2:")
	for i := 0; i < len(matrix2); i++ {
		fmt.Println(matrix2[i])
	}

	result := multiplication(matrix1, matrix2)
	fmt.Println("Результат умножения матриц:")
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
