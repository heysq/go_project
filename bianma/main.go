package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Away from keyboard"

	fmt.Println([]byte(message))

	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodedMessage)

	data, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(data))
}