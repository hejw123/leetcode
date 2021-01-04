package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i := 0 ; i < len(flowerbed) ; i = i + 2 {
		if flowerbed[i] == 0 {
			if i == len(flowerbed) - 1 || flowerbed[i + 1] == 0 {
				n --
			} else {
				i ++
			}
		}
	}
	return  n <= 0
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,0,1} , 2))
}
