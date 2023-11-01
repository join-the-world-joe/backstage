package main

var count int
var _encryption = false
var _userId = "1"
var _port = "10001"

func main() {
	if err := loop(_userId, _port, _encryption); err != nil {
		panic(err)
	}
}
