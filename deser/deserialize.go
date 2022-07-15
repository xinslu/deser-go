package deser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type deserialized map[string]interface{}

type string_serialized struct {
	input       string
	current     int
	deserialize deserialized
}

func deserialize(json string) deserialized {
	serialized := new(string_serialized)
	serialized.input = json
	serialized.current = 0
	serialized.deserialize = make(deserialized)
	serialized.cleanup()
	serialized.deserialize = serialized.parse_object()
	return serialized.deserialize
}

func (self *string_serialized) parse() interface{} {
	switch string(self.input[self.current]) {
	case "{":
		new_self := self
		new_self.deserialize = make(deserialized)
		return new_self.parse_object()
	case "\"":
		return self.parse_string()
	case ",":
		self.current++
	case "}":
		break
	case "[":
		return self.parse_array()
	default:
		if self.compare_contains("0123456789+-.e") {
			return self.parse_numbers()
		}
		result := self.get_string()
		if result != "" {
			resultlen := len(result)
			if result == "false" {
				self.current += resultlen
				return false
			} else if result == "true" {
				self.current += resultlen
				return true
			} else if result == "null" {
				self.current += resultlen
				return nil
			} else {
				panic("ERROR: Invalid string literal")
			}
		}
		panic("ERROR: Invalid value " + string(self.input[self.current]))
	}
	return nil
}

func (self *string_serialized) parse_object() deserialized {
	field := new(string)
	result := make(deserialized)
	for {
		switch string(self.input[self.current]) {
		case "{":
			self.current++
		case "\"":
			*field = self.parse_string()
		case ":":
			self.current++
			result[*field] = self.parse()
		case ",":
			self.current++
		case "}":
			break
		}
		if self.compare("}") {
			break
		}
	}
	return result
}

func (self *string_serialized) parse_array() []interface{} {
	result := make([]interface{}, 0)
	for {
		fmt.Println(string(self.input[self.current]))
		switch string(self.input[self.current]) {
		case "[":
			self.current++
		case "{":
			result = append(result, self.parse_object())
		case ",":
			self.current++
		case "]":
			break
		default:
			fmt.Println("char " + string(self.input[self.current]))
			result = append(result, self.parse())
			fmt.Println("char " + string(self.input[self.current]))
			fmt.Println(result...)
		}
		if self.compare("]") {
			fmt.Println("here" + string(self.input[self.current]))
			self.current++
			break
		}
	}
	return result
}

func (self *string_serialized) parse_string() string {
	string_object := new(string)
	self.match("\"")
	for !self.compare("\"") {
		*string_object += string(self.input[self.current])
		self.current++
	}
	self.current++
	return *string_object
}

func (self *string_serialized) parse_numbers() interface{} {
	int_object := new(string)
	isFloat := false
	for self.compare_contains("0123456789+-.e") {
		*int_object += string(self.input[self.current])
		if string(self.input[self.current]) == "." {
			isFloat = true
		}
		self.current++
	}
	if isFloat {
		result, err := strconv.ParseFloat(*int_object, 64)
		if err != nil {
			panic("ERROR: Cannot process flaot")
		}
		return result
	}
	result, err := strconv.Atoi(*int_object)
	if err != nil {
		panic("ERROR: Cannot process integer")
	}
	return result
}

func (self *string_serialized) get_string() string {
	current := self.current
	result := new(string)
	for unicode.IsLetter(rune(self.input[current])) {
		*result += string(self.input[current])
		current++
	}
	return *result
}

func (self *string_serialized) match(match string) {
	if string(self.input[self.current]) == match {
		self.current++
	} else {
		panic("ERROR: Expected a string")
	}
}

func (self *string_serialized) compare_contains(match string) bool {
	return strings.Contains(match, string(self.input[self.current]))
}

func (self *string_serialized) compare(match string) bool {
	return string(self.input[self.current]) == match
}

func (self *string_serialized) cleanup() {
	self.input = strings.Replace(self.input, "\n", "", -1)
	self.input = strings.Replace(self.input, " : ", ":", -1)
	self.input = strings.Replace(self.input, ": ", ":", -1)
	self.input = strings.Replace(self.input, " , ", ",", -1)
	self.input = strings.Replace(self.input, ", ", ",", -1)
	self.input = strings.Replace(self.input, " } ", "}", -1)
	self.input = strings.Replace(self.input, " }", "}", -1)
	self.input = strings.Replace(self.input, " { ", "{", -1)
	self.input = strings.Replace(self.input, "{ ", "{", -1)
}
