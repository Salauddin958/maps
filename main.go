package main

import (
	"fmt"
	"reflect"
)

func main() {

	var color1 map[string]string
	fmt.Println(color1)
	colors2 := make(map[string]string)
	fmt.Println(colors2)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	delete(colors, "red")
	printMap(colors)
	copyMap()
	compareMaps()
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func copyMap() {
	map_1 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	map2 := map[string]int{}
	for key, value := range map_1 {
		map2[value] = key
	}

	fmt.Println("Copied Map :", map2)
}

func compareMaps() {
	map1 := map[string]int{
		"x": 10,
		"y": 20,
		"z": 30,
	}
	map2 := map[string]int{
		"x": 10,
		"y": 20,
		"z": 30,
	}
	if reflect.DeepEqual(map1, map2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	map1 = map[string]int{
		"x": 10,
		"y": 20,
	}
	map2 = map[string]int{
		"x": 10,
		"y": 20,
		"z": 30,
	}
	if reflect.DeepEqual(map1, map2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
