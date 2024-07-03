package main

import (
	"bufio"
	"fmt"
	"os"
)

type Chars struct {
	index     int
	character rune
}

type Output struct {
	newLine string
	word    []Chars
}

func readBannerFile(bannerFile string) ([]string, error) {
	file, err := os.Open(bannerFile)
	if err != nil {
		return nil, err
	}
	var res []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res, nil
}

func getOutputSlice(s string) []Output {
	var outputStructSlice []Output
	var outputStruct Output
	var word []Chars
	var chars Chars
	var countChars int
	for i, ch := range s {
		if ch == '\n' {
			if countChars == len(s[i:i+countChars]) {
				outputStruct.word = word
				outputStructSlice = append(outputStructSlice, outputStruct)
				outputStruct = Output{}
				word = []Chars{}
				countChars = 0
			}
			outputStruct.newLine = "\n"
			outputStructSlice = append(outputStructSlice, outputStruct)
			outputStruct = Output{}
		} else {
			chars.index = i
			chars.character = ch
			word = append(word, chars)
			countChars++
		}
	}
	outputStruct.word = word
	outputStructSlice = append(outputStructSlice, outputStruct)
	return outputStructSlice
}

func indicesToColor(s1, s2 string) map[int]bool {
	indices := make(map[int]bool)
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if i+len(s2) <= len(s1) && s1[i:i+len(s2)] == s2 {
				indices = getIndicesToColor(i, i+len(s2), indices)
				i += len(s2) - 1
			}
		}
	}
	return indices
}

func getIndicesToColor(i, j int, indices map[int]bool) map[int]bool {
	for i < j {
		indices[i] = true
		i++
	}
	return indices
}

func getFinalOutput(s1, s2 string, lines []string) string {
	//var outputArt [][]string
	var output string
	//countLines := 0
	//res := ""
	var startIdx int
	var color1, reset, color2 string = "\033[0;32m", "\033[0m", ""
	outputStructSlice := getOutputSlice(s1)
	// fmt.Println(outputStructSlice)
	matchingRanges := indicesToColor(s1, s2)
	//fmt.Println(matchingRanges)
	for _, outputStruct := range outputStructSlice {
		if outputStruct.newLine == "\n" {
			output += "\n"
		} else {
			word := outputStruct.word
			if len(word) > 0 {
				//fmt.Println(wordMap)
				for k := 0; k < 8; k++ {
					for _, ch := range word {
						startIdx = int(ch.character-32)*9 + 1
						if len(color1) > 0 && matchingRanges[ch.index] {
							output += color1 + lines[startIdx+k] + reset
						} else if len(color2) > 0 && matchingRanges[ch.index] {
							output += color1 + lines[startIdx+k] + reset
						} else {
							output += lines[startIdx+k]
						}
					}
					output += "\n"
				}
			}

		}
	}
	return output
}
func main() {
	asciiLines, err := readBannerFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(getFinalOutput("hello\nthere\nthere", "re\nthe", asciiLines))
}
