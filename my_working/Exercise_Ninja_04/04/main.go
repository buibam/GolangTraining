package main

import (
	"fmt"
)

func main(){
	sl := []int{42,43,44,45,46,47,48,49,50,51}
	sl = append(sl, 52)
	fmt.Println(sl)
	sl = append(sl, 53, 54, 55)
	fmt.Println(sl)
	x := []int{56, 57, 58, 59}
	newsl := append(sl, x...)
	fmt.Println(newsl)
}
