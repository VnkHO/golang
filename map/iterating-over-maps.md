# Iterating over maps

```

  func main() {

    colors := map[string]string{
      "red":   "#ff0000",
      "green": "#00ff00",
      "blue":  "#0000ff",
      "white": "#ffffff",
    }

    fmt.Println(colors)
  }

  func printMap(c map[string]string) {
    for color, hex := rance c {

    }
  }

```

_c_ --> Argument name (colors)
_map[string]string_ --> Type of the map
_color_ --> Key for this step through the loop
_hex_ --> Valeu for this step through the loop

```

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

```

Hex code for blue is #0000ff
Hex code for white is #ffffff
Hex code for red is #ff0000
Hex code for green is #00ff00
