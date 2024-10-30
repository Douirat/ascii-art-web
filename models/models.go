package models

import (
	"bufio"
	"os"
	"strings"
)

// Handle user input:
func InputHandler(input string, data map[rune][]string) string {
	result := ""
	str := ""
	for i := 0; i < len(input); i++ {
		if rune(input[i]) == 13 && rune(input[i+1]) == 10 {
			result += Artify(str, data) + "\n"
			str = ""
			i ++
			continue
		}
		str += string(rune(input[i]))
	}
	if str != "" {
		result += Artify(str, data)
        return result
	}

	return result[:len(result)-1]
}



// ReadFileLineByLine reads a file and returns its content as a single string
func Readfile(filePath string) (map[rune][]string, error) {
	file, err := os.Open(filePath) // Open the file
	if err != nil {
		return nil, err
	}
	defer file.Close() // Ensure the file is closed when the function exits
	var sb strings.Builder // A string builder to store the content
	scanner := bufio.NewScanner(file)
	runesMap := make(map[rune][]string) 
	Rune := rune(32)
	index := 0
	for scanner.Scan() { // Read line by line
		sb.WriteString(scanner.Text()) // Add the line to the string builder
		sb.WriteString("\n")           // Add a newline character after each line
		index++
		runesMap[Rune] = append(runesMap[Rune], sb.String())
		if index != 0 && index % 9 == 0 {
			Rune ++
		}
		sb.Reset()
		
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return runesMap, nil // Return the entire content as a string
}

// form the ascii art representation of the user input:
func Artify(input string, mapData map[rune][]string) string {
	if input == "" {
		return ""
	}
	slc := [][]string{}
	result := ""
	for _, char := range input {
		slc = append(slc, mapData[char])
	}
	subSlc := slc[0]
	for x := 1; x < len(slc); x++ {
		for y := 0; y < len(slc[x]); y++ {
			subSlc[y] = subSlc[y][:len(subSlc[y])-1] + slc[x][y]
		}
	}
	for _, char := range subSlc {
		result += string(char)
	}
	return result[1 : len(result)-1]
}



