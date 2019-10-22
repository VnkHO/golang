# Manipulating maps

```

func main() {

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }

	//var colors map[string]string

  colors := make(map[string]string)

	colors["white"] = "#ffffff"
	delete(colors, "white")


	fmt.Println(colors)
}

```

Whenever we declare a new variable and if we do not assign an actual value to it; GO will initialize it with it's _zero value_

We do not get **dot .** syntax with maps.
We have to use **square brackets []** because the key that we provide has to be of the appropriate type.
