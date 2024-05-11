/*
Сумма обратных чисел
Дано два натуральных числа, не заканчивающиеся на 00. Замените каждое число на обратное и вычислите их сумму. 
Например, дается два числа 624624 и 10241024. Каждое заменяем на обратное, то есть 624⇒426,1024⇒4201.624⇒426,1024⇒4201. 
Затем находим их сумму: 426+4201=4627.426+4201=4627.
 
Входные данные
Вводятся два натуральных числа, каждый из которых не превосходит 108108. Гарантируется, что каждое число не оканчивается на 00.

Выходные данные
Выведите сумму обратных чисел.
*/
package main

import ("fmt"; "strconv")

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	su := reverseValue(a) + reverseValue(b)
	fmt.Println(su)

}

func reverseValue(x int) int {
	var s string = ""
	for x > 0 {
		s += strconv.Itoa(x % 10)
		x = x / 10
	}
	z, _ := strconv.Atoi(s)
	return z
}
// Решено