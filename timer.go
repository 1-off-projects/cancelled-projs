package main

import (
	"fmt"
	"time"
)

func main() {
	hr, err := fmt.Scanln("enter hours:")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	hr = (hr * 3600000000000)
	min, err := fmt.Scanln("enter minutes:")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	min = (min * 60000000000)
	sec, err := fmt.Scanln("enter seconds:")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	sec = (sec * 100000000)
	time.Sleep(time.Duration(hr + min + sec))
	fmt.Print("pong")
}
