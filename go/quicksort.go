package main

import "fmt"

func QuickSort(array []int) {
	Quick(array, 0, len(array)-1)
}
 
func Quick(array []int, l, r int) {
	if l >= r {
		return
	}
	point := Partition(array, l, r)
	Quick(array, l, point-1)
	Quick(array, point+1, r)
}
 

func Partition(array []int, l, r int) int {
	point := array[r]
	i := l
	for j := l; j < r; j++ {
		if array[j] < point {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[r] = array[r], array[i]
	return i
}

func main(){
	fmt.Println("Please input 5 numbers:");
	var a  [] int
	var n int
	for i := 0; i < 5; i++{
		fmt.Scan(&n)
		a = append(a, n)
	}	
	QuickSort(a)
	for _, v := range a {
		fmt.Printf("%5d",  v)
	}
}