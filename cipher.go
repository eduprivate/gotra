package main 

import (
	"bufio"
	"fmt" 
	"errors"
	"io"
	"strings"
	"os"		
)

func cipher(phrase string, key int32) string {
	content := ""
	for i := 0; i < len(phrase); i++ {
		var c = []rune(phrase)[i]
		c = c + key
		fmt.Printf("%c", c)
		content = content + string(c)
	}
	return content
}

func decipher(phrase string, key int32) string {
	content := ""
	for i := 0; i < len(phrase); i++ {
		var c = []rune(phrase)[i]
		c = c + key
		fmt.Printf("%c", c)
		content = content + string(c)
	}
	return content
}

func loadFile(fname string) (string, error) {
	if fname == "" {
		return "", errors.New("Decripted file name cannot be emtpy.")
	}

	// attempt to load dict file.
	file, err := os.Open(fname)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	content := ""
	for _, value := range lines {
		content = content + value
	} 
	fmt.Println(content)
	return content, scanner.Err()
}

func writeFile(fname, content string) error {
	fo, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(content))
	if err != nil {
		return err
	}

	return nil
}

func decriptFile(file_origen, file_destine string, key int32) error {
	content, err := loadFile(file_origen)
	
	if err != nil {
		return err
	}

	file_content := decipher(content, key)
	writeFile(file_destine, file_content)
	return nil
}

func criptFile(file_origen, file_destine string, key int32) error {
	content, err := loadFile(file_origen)

	if err != nil {
		return err
	}

	file_content := cipher(content, key)
	writeFile(file_destine, file_content)
	return nil
}

func main() {
	// fmt.Println("Caesar Cipher: ")
	// fmt.Println("Cipher: ")
	// cipher("Olha que legal", 3);
	// fmt.Println("Decipher: ")
	// decipher("Rokd#txh#ohjdo",3);

	criptFile("./textin.txt", "./textout.txt", 3)


}