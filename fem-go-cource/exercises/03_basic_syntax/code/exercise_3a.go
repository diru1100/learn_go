package main

import "fmt"

func main() {

	sample_sentence := " lorem ipsum this is a fake sentence :p"

	for index, value := range sample_sentence {
		if index%2 == 0 {
			fmt.Println(" This is the index =", index, "and value(in string)=", string(value))
		}
		
		


	}
}