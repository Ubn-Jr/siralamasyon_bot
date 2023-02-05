// SIRALAMASYON BOT
// NOTLAR:
// * Her fonksiyonun en az üç testi olmalı
package main

// Bir string değeri alıp bu değerin
// liste olup olmadığını kontrol eden fonksiyon
// Liste süslü parantezle başlamalı ve bitmeli
// Liste elemanları virgülle ayrılmalı
// Kapatan süslü parantezden önceki son karakter ve
// Açan süslü parantezden sonraki değer virgül olmamalı
func stringListeMi(discordMesaji string) bool {
	// TODO: Bu fonksiyonun içi doldurulacak : Beyzanur
	if discordMesaji[0] != '{' || discordMesaji[len(discordMesaji)-1] != '}' {
		return false
	} else {
		return true
	}
}

// Gelen string degerini listeye cevir
func stringiListeyeCevir(discordMesaji string) []int {
	// TODO: Bu fonksiyonun içi doldurulacak : Hatice
	return []int{}
}

func dizidekiDegerlerSayiMi(dizi []string) bool {
	// TODO: Bu fonksiyonun içi doldurulacak : Hatice
	return true
}

// Kullanıcıya hata mesajı göster
func hataMesajiGoster(hataMesaji string) {
	// TODO: Bu fonksiyonun içi doldurulacak : Beyzanur
}

func main() {
	// stringListeMi fonksiyonunu test et
	// var listeMi1 = stringListeMi("12341413123123") // false

	// stringiListeyeCevir fonksiyonunu test et

	// dizidekiDegerlerSayiMi fonksiyonunu test et

	// hataMesajiGoster fonksiyonunu test et
}
