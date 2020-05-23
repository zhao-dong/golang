package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	m := 0
	for m < 10 {
		fmt.Println(m)
		m++
	}

	/*
		a := [3]int{1, 2, 3}
		for _, v := range a { //复制一份a遍历[1, 2, 3]
			v += 100 //v是复制对象中的值，不会改变a数组元素的值
		}
		fmt.Println(a) //1 2 3
	*/

	/*
		a := [3]int{1, 2, 3}
		for i, v := range a { //i,v从a复制的对象里提取出
			if i == 0 {
				a[1], a[2] = 200, 300
				fmt.Println(a) //输出[1 200 300]
			}
			a[i] = v + 100 //v是复制对象里的元素[1, 2, 3]
		}
		fmt.Println(a) //输出[101, 102, 103]
	*/

	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)
}
