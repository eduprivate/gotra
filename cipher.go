package main 

import "fmt"

func cipher(phrase string, key int) {
	for i := 0; i < len(phrase); i++ {
		var c = []rune(phrase)[i]
		for j := 0; j < key; j++ {
			c = c + 1
		}
		fmt.Printf("%c", c)
	}
}

func decipher(phrase string, key int) {
	for i := 0; i < len(phrase); i++ {
		var c = []rune(phrase)[i]
		for j := 0; j < key; j++ {
			c = c - 1
		}
		fmt.Printf("%c", c)
	}
}

func main() {
	fmt.Println("Caesar Cipher: ")
	fmt.Println("Cipher: ")
	cipher("Olha que legal", 3);
	fmt.Println("Decipher: ")
	decipher("Rokd#txh#ohjdo",3);
}