package main

import (
	"fmt"
)

// type class struct{
// 	X int
// 	Y int
// }
// func(c *class) methood(){

// }
// type point struct {
// 	X int

// 	Y int
// }

func main() {
	list := make([]int, 5, 10)
	fmt.Println(len(list), cap(list))

}

// func pointmapp(){
// 	pointsMap := map[string]point{}
// 	// OtherpointMap := make(map[int]point)

// 	pointsMap["adad"] = point{X: 1, Y: 45}

// 	fmt.Println(pointsMap["adad"])
// }
// func rangefor()  {
// 	arr := []int{1, 2, 3, 4, 5, 6}

// 	for _, l := range arr {

// 		if l == 4 {
// 			fmt.Println(l)
// 		} else {
// 			fmt.Println("это не чотыре ")
// 		}
// 	}

// func arrays()  {
// 	var a [2] string
// 	a[0]="privet"
// 	a[1]="mir"
//     fmt.Println(a)

// 	numbers:=[...]int{1,3,6,7,8,8}

// }

// func strut()  {
// 	p1:=class{
// 		X: 10,
// 		Y: 20,
// 	}
// 	fmt.Println(p1)
// }
// func aboba() {
// 	a := "Gogag"
// 	b := 132
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	p := &a
// 	fmt.Println(*p)

// }
