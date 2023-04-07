package main
import "fmt"
// buyukten kucuge siralar 
func selectionSortBigToSmall(discordMesaji []int) {
	for i := 0; i < len(discordMesaji); i++ {
		smallNumber := i
		for j := i + 1; j < len(discordMesaji); j++ {
			if discordMesaji[j] > discordMesaji[smallNumber] {
				smallNumber = j
			}
		}
		discordMesaji[i], discordMesaji[smallNumber] = discordMesaji[smallNumber], discordMesaji[i]
	}
	fmt.Println(discordMesaji)
}

func main() {
	arr := []int{522, 5, 1, 8, 2, 6, 9, 62, 12}
	selectionSortBigToSmall(arr)
}
