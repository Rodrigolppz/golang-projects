package main

func main() {

	subject := "stone"
	candidates := []string{"snow", "stone", "tones", "banana", "tons", "notes", "Seton"}

	for i, _ := range candidates {

		if len(subject) == len(candidates[i]) {
			return true

		}

	}

}
