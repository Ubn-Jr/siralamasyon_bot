package main 

import "fmt"



func selectionSortSmallToBig(discordMesaji []int) {
	//Kucukten buyuge siralar 
	for i:= 0; i <len(discordMesaji); i ++ {
		smallNumber := i
		for j := i + 1 ; j < len(discordMesaji); j++ {
			if discordMesaji[j] < discordMesaji[smallNumber]{
				smallNumber = j 
			}	
		} 
		discordMesaji[i], discordMesaji[smallNumber] = discordMesaji[smallNumber], discordMesaji[i]
	}
  fmt.Println(discordMesaji)
}
func main(){

var arr []int = []int{522,5,1,8,2,6,9,62,12}

selectionSortSmallToBig(arr)
}