package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Converting Hex to base64")
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	p, err := hex.DecodeString(hexStr)
	if err != nil {
		log.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(p)
	fmt.Println(base64Str)

}
