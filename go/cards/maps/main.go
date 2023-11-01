package main

import "fmt"

func main() {
	colors := map[string]string{
		"black": "#144234",
		"white": "#000000",
	}

	// Ou
	colors["green"] = "ffffff"
	printMap(colors)

}

func printMap(c map[string]string){
	for color, hex := range c {
		//green : ffffff
		//black : #144234
		//white : #000000
		fmt.Println(":",  hex)

	}
}