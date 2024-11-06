package main

import (
	"fmt"
	"math/rand"
)

func randomAccess(a [5]int) int {
	// 在区间 [0, nums.length) 中随机抽取一个数字
	randomIndex := rand.Intn(len(a))
	return a[randomIndex]
}
func insert(a *[5]int, index int, value int) {
	for i := len(a) - 1; i > index; i-- {
		a[i] = a[i-1]
	}
	a[index] = value

}
func delete(a *[5]int, index int) {
	for i := index; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
}
func traverse(a *[5]int) {
	count := 0
	// 通过索引遍历数组
	for i := 0; i < len(a); i++ {
		count += a[i]
	}
	count = 0
	// 直接遍历数组元素
	for _, num := range a {
		count += num
	}

}
func find(a *[5]int, target int) int {
	index := -1 //返回-1表示没找到
	for i := 0; i < len(a); i++ {
		if a[i] == target {
			index = i
			break
		}
	}
	return index
}

func main() {
	//定义数组
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
}
