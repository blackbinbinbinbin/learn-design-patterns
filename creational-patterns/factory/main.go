package main

import (
	"fmt"
)

func main() {
	win, _ := GetGuiFactory("window")
	mac, _ := GetGuiFactory("mac")

	win.setButton("this is a windows button")
	mac.setButton("this is a mac button")

	fmt.Println(fmt.Sprintf("Button: %s", win.getButton()))
	fmt.Println(fmt.Sprintf("Button: %s", mac.getButton()))
}
