package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["NY"] = "New York"
	states["NJ"] = "New Jersey"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "WA")
	fmt.Println(states)

	states["TX"] = "Texas"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
