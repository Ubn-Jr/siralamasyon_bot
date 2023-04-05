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
   

     chars := []byte(discordMesaji)
 
     for i := 0; i < len(chars); i++ {
         char := string(chars[i])
         fmt.Println(char)
     }
     
     if chars == "" {
    fmt.Println("Hata: Boş bir mesaj verildi!")
    return false
    } else {
        return true
    }

 
 }

