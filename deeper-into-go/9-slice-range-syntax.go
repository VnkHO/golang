/*

	Slices are zero-indexed

						0					1					2				3
	fruits = "apple" "banana" "grape" "orange"

	fruits[0] = "apple"
	fruits[3] = "orange"

	fruits[startIndexIncluding: upToNotIncluding]
	fruits[0: 2] -> "apple" "banana"
	fruits[:2] -> "apple" "banana" // to the very beginning
	fruits[2:] -> "grape" "orange" // to the last end

*/