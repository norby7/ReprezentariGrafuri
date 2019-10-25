package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func citireMatriceAdiacenta() [][]int {
	file, err := ioutil.ReadFile("matriceAdiacenta.txt")

	if err != nil {
		fmt.Println("Eroare citire matrice de adiacenta.")
	}

	lines := strings.Split(string(file), "\r")
	n := len(lines)

	matrix := make([][]int, n)

	for i, line := range lines {
		matrix[i] = make([]int, n)

		vals := strings.Split(strings.TrimPrefix(line, "\n"), " ")

		for j, value := range vals {
			//fmt.Println(fmt.Sprintf("%d %d \n",i,j))
			matrix[i][j], _ = strconv.Atoi(value)
		}
	}

	return matrix
}

func citireMatriceIncidenta() [][]int {
	file, err := ioutil.ReadFile("matriceIncidenta.txt")

	if err != nil {
		fmt.Println("Eroare citire matrice de incidenta.")
	}

	lines := strings.Split(string(file), "\n")
	n := len(lines)

	matrix := make([][]int, n)

	for i, line := range lines {
		vals := strings.Split(line, " ")

		matrix[i] = make([]int, len(vals))

		for j, value := range vals {
			matrix[i][j], _ = strconv.Atoi(value)
		}
	}

	return matrix
}

func adiacentaToIncidenta(matrix [][]int) [][]int {
	nodes := 0
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				nodes++
			}
		}
	}

	matrixI := make([][]int, n)
	nodeIndex := 0

	for i := 0;i< n;i++{
		matrixI[i] = make([]int, nodes)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1{
				matrixI[i][nodeIndex] = 1
				matrixI[j][nodeIndex] = -1
				nodeIndex++
			}
		}
	}

	return matrixI
}

func main() {
	matrix := citireMatriceAdiacenta()
	fmt.Println(matrix)

	matrixI := adiacentaToIncidenta(matrix)
	fmt.Println(matrixI)
}
