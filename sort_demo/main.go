package main

import (
	"fmt"
	"sort"
)

func main() {
	//GuessingGame()
	//SortIntSliceDemo()
	//SortInterfaceSliceDemo()
	SortInterfaceSliceStableDemo()
}

// GuessingGame 猜数游戏
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		_, _ = fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func SortIntSliceDemo() {
	a := []int{2, 3, 4, 200, 100, 21, 234, 56}
	x := 21

	sort.Ints(a)
	index := sort.Search(len(a), func(i int) bool { return a[i] >= x }) // 查找元素

	if index < len(a) && a[index] == x {
		fmt.Printf("found %d at index %d in %v\n", x, index, a)
	} else {
		fmt.Printf("%d not found in %v,index:%d\n", x, a, index)
	}
}

func SortInterfaceSliceDemo() {
	people := []struct {
		Name string
		age  int
	}{
		{"张三", 16},
		{"李四", 86},
		{"王五", 59},
		{"赵六", 37},
		{"周七", 14},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].age <= people[j].age
	})
	fmt.Println("Sort by age:", people)
}

func SortInterfaceSliceStableDemo() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age >= people[j].Age
	}) // 按年龄降序排序
	fmt.Println("Sort by age:", people)
}
