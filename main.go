package main

import (
	"fmt"
)

func main() {
	input := "{ \"hello\" : -1 , \"bruh\" : \"here\", \"objectbruh\" : {\"here\": -1.01, \"objectbruh\" : {\"here\": false, \"second\": [1,2,3,\"here\"]} } }"
	fmt.Println(deser.deserialize(input))
}
