package main

var (
	charValid = map[string]bool{
		"<": true,
		">": true,
		"{": true,
		"}": true,
		"[": true,
		"]": true,
	}
	openCloseCharValid = map[string]string{
		"<": ">",
		">": "<",
		"{": "}",
		"}": "{",
		"[": "]",
		"]": "[",
	}
)

func main() {
	// <>{}[]
}
