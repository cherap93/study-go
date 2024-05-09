/*
Дан список. Выведите те его элементы, которые встречаются в списке только один раз.

Входные данные
Сначала задано число NN — количество элементов в массиве (1≤N≤100001≤N≤10000). Далее через пробел записаны NN чисел — элементы массива.
Массив состоит из целых чисел, каждое из которых по модулю не превышает 1000010000.

Выходные данные
Необходимо вывести на одной строке через пробел элементы, которые встречаются в списке только один раз.
Элементы нужно выводить в том порядке, в котором они встречаются в списке.
*/
package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)

	sli1 := make([]int, size)

	flag := "alone"

	for i := 0; i < size; i++ {
		fmt.Scan(&sli1[i])
	}

	for i := 0; i < size; i++ {
		flag = "alone"
		for j := 0; j < size; j++ {
			if j == i {
				continue
			}
			if sli1[i] == sli1[j] {
				flag = "multi"
				break
			}
		}
		if flag == "alone" {
			fmt.Print(sli1[i], " ")
		}
	}

}

// решено
