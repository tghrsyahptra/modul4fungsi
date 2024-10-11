package main

import "fmt"

func main() {
	var int1, int2, int3, int4, int5 int
	var char1, char2, char3 rune

	fmt.Println("Masukan 5 buah nilai antara 32 - 127 : ")
	fmt.Scan(&int1, &int2, &int3, &int4, &int5)
	fmt.Printf("%c %c %c %c %c\n", int1, int2, int3, int4, int5)

	fmt.Scanln()
	
	fmt.Println("Masukkan 3 karakter:")
	fmt.Scanf("%c%c%c", &char1, &char2, &char3) // Perhatikan spasi antara %c untuk input spasi
	
	fmt.Printf("%c%c%c\n", char1, char2, char3)
	
}
