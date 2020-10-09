package main

import "fmt"

/**
func add(x int, y int) int{} 等价于
int add(int x, int y){}
*/
func add(x int, y int) int {
	return x+y
}

func main(){
	fmt.Println(add(42,13))
}