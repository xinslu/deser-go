package main

import (
    "fmt"
)

type deserialized map[string]interface{}

type string_serialized struct {
    input string
    current int
}


func main() {
    input := "{\"hello\" : 1}"
    fmt.Print(deserialize(input))
}


func deserialize(json string) deserialized{
    return deserialized {
        "bruh":2,
    }
}

func (self *string_serialized) parse_string(string_input string) string {
    string_object := new(string)
    self.match("\"")
    for !self.compare("\"") {
        *string_object += string(self.input[self.current])
        self.current++
    }
    return *string_object
}

func (self *string_serialized) match(match string) {
    if string(self.input[self.current]) == match {
        self.current++
    } else {
        panic("Expected a string")
    }
}

func (self *string_serialized) compare(match string) bool{
    return string(self.input[self.current]) == match
}
