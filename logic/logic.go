package logic

import (
	"bufio"
	"os"
	"strings"
)

// Read data from a file:
// ReadFileLineByLine reads a file and returns its content as a single string
func Readfile(filePath string) ([]string, error) {
	file, err := os.Open(filePath) // Open the file
	if err != nil {
		return nil, err
	}
	defer file.Close() // Ensure the file is closed when the function exits
	var slc []string
	var sb strings.Builder // A string builder to store the content
	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // Read line by line
		sb.WriteString(scanner.Text()) // Add the line to the string builder
		sb.WriteString("\n")           // Add a newline character after each line
		slc = append(slc, sb.String())
		sb.Reset()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return slc, nil // Return the entire content as a string
}

func HashCharacter(r rune, data []string) []string {
	start_index := len(data) - ((127 - int(r)) * 9)
	end_index := start_index + 9
	return data[start_index:end_index]
}

// form the ascii art representation of the user input:
func Artify(input string, data []string) string {
	if input == "" {
		return ""
	}
	slc := [][]string{}
	result := ""
	for _, char := range input {
		slc = append(slc, HashCharacter(char, data))
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

// Handle user input:
func InputHandler(input string, data []string) string {
	result := ""
	str := ""
	for i := 0; i < len(input); i++ {
		if rune(input[i]) == 92 && rune(input[i+1]) == 110 {
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
