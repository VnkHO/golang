package main

import "fmt"

func main() {

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }

	// var colors map[string]string
	colors := make(map[string]string)

	colors["white"] = "#ffffff"

	delete(colors, "white")

	fmt.Println(colors)
}
