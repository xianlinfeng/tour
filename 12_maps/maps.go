package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)

	var n = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(n)

	p := make(map[string]int)
	p["Answer"] = 42

	// modify a map
	p["Answer"] = 21

	//delte a key
	delete(p, "Answer")

	// check a value is exist or not
	v, ok := p["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
