package main

type FlagDetails struct {
	description string
	callback    func(string) (int, error)
}

func Flags() map[string]FlagDetails {
	return map[string]FlagDetails{
		"c": {
			description: "number of bytes in the file.",
			callback:    numBytes,
		},
		"l": {
			description: "Number of lines in the file.",
			callback:    numLines,
		},
		"w": {
			description: "Number of words in the file.",
			callback:    numWords,
		},
		"m": {
			description: "Number of special characters in the file.",
			callback:    numChar,
		},
	}
}
