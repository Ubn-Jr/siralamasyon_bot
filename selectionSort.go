package main 

import "fmt"

func main(){

var arr []int = []int{522,5,1,8,2,6,9,62,12}

smallNumber := 0

for i:= 0; i <len(arr); i ++ {
	smallNumber = i
	for j := i + 1 ; j < len(arr); j++ {
		if arr[j] < arr[smallNumber]{
			smallNumber = j 
		}	
	} 
    arr[i], arr[smallNumber] = arr[smallNumber], arr[i]
}
fmt.Println(arr)

}