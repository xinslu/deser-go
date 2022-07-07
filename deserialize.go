package main

import (
	"fmt"
	"strings"
)

type deserialized map[string]interface{}

type string_serialized struct {
	input       string
	current     int
	deserialize deserialized
}

func main() {
	input := "{ \"hello\" : 1 , \"bruh\" : 2 }"
	fmt.Print(deserialize(input))
}

func deserialize(json string) deserialized {
	json = strings.Replace(json, " : ", ":", -1)
	json = strings.Replace(json, " }", "}", -1)
	serialized := new(string_serialized)
	serialized.input = json
	serialized.current = 0
	serialized.deserialize = make(deserialized)
	serialized.cleanup()
	fmt.Println(serialized.input)
	fmt.Println("here")
	for {
		fmt.Println(string(serialized.input[serialized.current]))
		switch string(serialized.input[serialized.current]) {
		case "{":
			serialized.current++
		case "\"":
			field := serialized.parse_string()
			serialized.current += 2
			serialized.deserialize[field] = serialized.parse_integer()
		case ",":
			serialized.current++
		case "}":
			break
		}
		if serialized.compare("}") {
			break
		}
	}
	return serialized.deserialize
}

func (self *string_serialized) parse_object() {

}

func (self *string_serialized) cleanup() {
	self.input = strings.Replace(self.input, "\n", "", -1)
	self.input = strings.Replace(self.input, " : ", ":", -1)
	self.input = strings.Replace(self.input, " , ", ",", -1)
	self.input = strings.Replace(self.input, ", ", ",", -1)
	self.input = strings.Replace(self.input, " } ", "}", -1)
	self.input = strings.Replace(self.input, " }", "}", -1)
	self.input = strings.Replace(self.input, " { ", "{", -1)
	self.input = strings.Replace(self.input, "{ ", "{", -1)
}

func (self *string_serialized) parse_string() string {
	string_object := new(string)
	self.match("\"")
	for !self.compare("\"") {
		*string_object += string(self.input[self.current])
		self.current++
	}
	return *string_object
}

func (self *string_serialized) parse_integer() int {
	int_object := new(int)
	for self.compare_contains("123456789") {
		*int_object = *int_object*10 + (int(self.input[self.current]) - 48)
		self.current++
	}
	return *int_object
}

func (self *string_serialized) match(match string) {
	if string(self.input[self.current]) == match {
		self.current++
	} else {
		panic("Expected a string")
	}
}

func (self *string_serialized) compare_contains(match string) bool {
	return strings.Contains(match, string(self.input[self.current]))
}

func (self *string_serialized) compare(match string) bool {
	return string(self.input[self.current]) == match
}
