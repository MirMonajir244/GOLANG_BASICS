package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var freq [128]int
	var str string
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string: ")
	sc.Scan()
	str = sc.Text()

	for i := 0; i < len(str); i++ {
		count := 0
		index := int(str[i])
		if freq[index] == 0 {
			freq[index]++
			for j := i; j < len(str); j++ {
				if str[j] == str[i] {
					count++
				}
			}
			fmt.Printf("%s%d ", string(str[i]), count)
		}
	}
}
