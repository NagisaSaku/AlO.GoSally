package main

 import (
     "fmt"
 )

func bubble(num []int) []int {
	flag := true
	for i:= len(num) -1 ; i >= 0 ; i--{
		for j:= 0 ; j < i ; j++{
			if num[j] > num[ j + 1 ]{
				Swap(num , j , j + 1)
				flag = false
			}
		} 
		if flag {
			break
		}
	}
	return num
}

func Swap(num []int,i int,j int)  {
	temp := num[i]
	num[i] = num[j]
	num[j] = temp
}

func main()  {
	var numArray = []int{2,3,43,1,12}
	var result =  bubble(numArray)
	for i:= len(result) -1 ; i >= 0 ; i--{
		fmt.Println(result[i])
	}
}