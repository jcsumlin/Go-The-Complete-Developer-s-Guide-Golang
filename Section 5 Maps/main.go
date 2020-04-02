package main

import (
	"fmt"
)

func main() {

	// Maps are also similar to Dicts in Python and an Object in Javascript
	// Both the keys and Values are statically typed

	// all the keys and values are strings in the colors map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors) //map[red:#ff0000]

	// Method 2 of declaring a map: if we want to figure out later on what values we want to add to the map
	var colors2 map[string]string
	fmt.Println(colors2) //map[]

	//Method 3:
	colors3 := make(map[string]string)
	fmt.Println(colors3)
	colors3["key"] = "value"
	colors3["white"] = "#ffffff"
	fmt.Println(colors3)          //map[key:value white:#ffffff]
	fmt.Println(colors3["white"]) // #ffffff
	// We use the square brace syntax with maps to get a specific value
	delete(colors3, "key")
	fmt.Println(colors3) //map[white:#ffffff]

	// Itterating over a map
	printMap(colors)
	//red #ff0000
	//green #4bf745
	//white #ffffff

	//How is a map different from a struct?
	// Annotation 2020-04-02 103819.png
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("%s %s\n", key, value)
	}
}
