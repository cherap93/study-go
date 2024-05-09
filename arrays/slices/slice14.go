/*
Таблица умножения

Даны два числа nn и mm. Создайте двумерный массив A[n,m]A[n,m], заполните его таблицей умножения A[i,j]=i⋅jA[i,j]=i⋅j и выведите на экран.

Входные данные
Программа получает на вход строку, в которой через пробел даны два числа nn и mm – количество строк и столбцов, соответственно.

Выходные данные
Программа должна вывести полученный массив. Числа разделяйте одним пробелом.
*/
package main

import "fmt"

func main() {
	var size1, size2 int
	fmt.Scan(&size1, &size2)

	sli := make([][]int, size1)
	for i := 0; i < size1; i++ {
		sli[i] = make([]int, size2)
	}

	for i := 0; i < size1; i++ {
		for j := 0; j < size2; j++ {
			sli[i][j] = (i + 1) * (j + 1)
			fmt.Print(sli[i][j], " ")
		}
		fmt.Println()
	}

}

// решено
