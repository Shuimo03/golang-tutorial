package main
import "fmt"

func split(sum int)(x,y int){
	x = sum*4/9 // sum*4/9 取整
	y = sum-x // sum-x
	return
}

func main(){
	fmt.Println(split(200))
}