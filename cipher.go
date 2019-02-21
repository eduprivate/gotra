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

func decipher(phase string, key int) {

}

func main() {
	cipher("Olha que legal", 3);
}