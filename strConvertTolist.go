package main 
import ( 
	"fmt"
	"strconv"
)

func main() {
  
 var discordMesaji  string = "deneme mesaji"
 stringiListeyeCevir(discordMesaji)
}
// Gelen string degerini listeye cevir
func stringiListeyeCevir(discordMesaji string) []int {
    // Discord mesaji parametresine
     chars := []byte(discordMesaji)
 
     for i := 0; i < len(chars); i++ {
         char := string(chars[i])
         fmt.Println(char)
     }
     
 
 }

