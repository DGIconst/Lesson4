package main

import "fmt"

func main() {
	var arr = make([]int, 0, 1000000)
	var a, b, c, d, e, f, g, h int
	fmt.Println("Добавте любые целые числа")
	fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h)
	arr = append(arr, a, b, c, d, e, f, g, h)
	InsertionSort(arr)
	fmt.Println(arr)
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j >= 1 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
}

// не смог разобраться как добовлять любое количество чисел в массив с панели ( а то у меня получается только то скольок переменных указано)
