package main

import "fmt"

func basicExpresionlessSwitch(val int) bool {
	var result bool
	switch x := 10; { // expressionless switch
	case val >= x:
		result = true
	case val < x:
		result = false
	}
	return result
}

func typeSwitch(i interface{}) string {
	switch i.(type) {
	case int:
		return "it's an integer!"
	case string:
		return "it's a string"
	default:
		return "type is unknown"
	}
}

func main() {
	fmt.Println(basicExpresionlessSwitch(10))
	fmt.Println(typeSwitch(10))
	fmt.Println(typeSwitch("aha"))
}
