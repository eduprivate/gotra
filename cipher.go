package main 

import (
	"bufio"
	"fmt" 
	"errors"
	"os"		
)

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

func loadFile(fname string) ([]string, error) {
	if fname == "" {
		return nil, errors.New("Decripted file name cannot be emtpy.")
	}

	// attempt to load dict file.
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeFile(fname, content string) {

}

func main() {
	fmt.Println("Caesar Cipher: ")
	fmt.Println("Cipher: ")
	cipher("Olha que legal", 3);
	fmt.Println("Decipher: ")
	decipher("Rokd#txh#ohjdo",3);
}