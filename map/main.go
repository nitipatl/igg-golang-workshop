package main

import "fmt"

func main() {

	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#00FF00",
	// 	"black": "#000000",
	// }
	// fmt.Println(colors)

	colors := make(map[string]string)
	colors["red"] = "#FF0000"
	colors["black"] = "#000000"
	updateMap(colors, "black", "#ffffff")
	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hexcode := range colors {
		fmt.Println(color, hexcode)
	}
}

func updateMap(colors map[string]string, color string, code string) {
	colors[color] = code
}
