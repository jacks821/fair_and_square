package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve(a []string, palins []int64) int {
	var i int64
	var n int64
	i, _ = strconv.ParseInt(a[0], 10, 64)
	n, _ = strconv.ParseInt(a[1], 10, 64)
	return PalinsInLine(i, n, palins)
}

func PalinsInLine(i int64, n int64, palins []int64) int {
	cases := 0
	for _, palindrome := range palins {
		if palindrome >= i && palindrome <= n {
			cases += 1
		}
	}
	return cases
}

func GeneratePalindromesWithLength(l int) []int64 {
	var palindromes []int64
	if l == 1 {
		for i := int64(1); i < 10; i++ {
			palindromes = append(palindromes, i)
		}
		return palindromes
	}
	if l%2 == 1 {
		half_length := int64((l - 1) / 2)
		for x := int64(0); x < int64(10); x++ {
			for y := int64(math.Pow(10, float64(half_length-1))); y < int64(math.Pow(10, float64(half_length))); y++ {
				ystring := strconv.FormatInt(y, 10)
				xstring := strconv.FormatInt(x, 10)
				palstring := ystring + xstring + Reverse(ystring)
				palindrome, _ := strconv.ParseInt(palstring, 10, 64)
				palindromes = append(palindromes, palindrome)
			}
		}
	} else {
		half_length := int64(l / 2)
		for x := int64(math.Pow(10, float64(half_length-1))); x < int64(math.Pow(10, float64(half_length))); x++ {
			xstring := strconv.FormatInt(x, 10)
			palstring := xstring + Reverse(xstring)
			palindrome, _ := strconv.ParseInt(palstring, 10, 64)
			palindromes = append(palindromes, palindrome)
		}
	}
	return palindromes
}

func GeneratePalindromes(min int64, max int64) []int64 {
	max_string := strconv.FormatInt(max, 10)
	min_string := strconv.FormatInt(min, 10)
	min_len := len(min_string)
	max_len := len(max_string)
	var palindromes []int64
	for l := min_len; l <= max_len; l++ {
		for _, x := range GeneratePalindromesWithLength(l) {
			if x >= min && x <= max {
				palindromes = append(palindromes, x)
			}
		}
	}
	return palindromes
}

func isPalindrome(s int64) bool {
	i := strconv.FormatInt(s, 10)
	reversed, _ := strconv.ParseInt(Reverse(i), 10, 64)
	return s == reversed
}

func GrabLines(args string) []string {
	var lines []string
	file, err := os.Open(args)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func isSquare(s int64) bool {
	root := int64(math.Sqrt(float64(s)))
	return (root * root) == s
}

func Reverse(s string) string {
	arr := strings.Split(s, "")
	var newarr []string
	for i := len(arr) - 1; i >= 0; i-- {
		newarr = append(newarr, arr[i])
	}
	newstring := strings.Join(newarr, "")
	return newstring
}

func GeneratePalins(lines []string) []int64 {
	var palins []int64
	var init int64
	var ending int64
	for index, line := range lines {
		arr := strings.Split(line, " ")
		i, _ := strconv.ParseInt(arr[0], 10, 64)
		n, _ := strconv.ParseInt(arr[1], 10, 64)
		if index == 0 {
			init = i
			ending = n
		} else {
			if i < init {
				init = i
			}
			if n > ending {
				ending = n
			}
		}
	}
	for _, palin := range GeneratePalindromes(init, ending) {
		if isSquare(palin) {
			root := int64(math.Sqrt(float64(palin)))
			if isPalindrome(root) {
				palins = append(palins, palin)
			}
		}
	}
	return palins
}

func main() {
	argsWithoutProgram := os.Args[1]
	lines := GrabLines(argsWithoutProgram)
	cases, _ := strconv.Atoi(lines[0])
	palins := GeneratePalins(lines[1:])
	for i := 1; i <= cases; i++ {
		line := strings.Split(lines[i], " ")
		fmt.Printf("Case #%d: %d\n", i, Solve(line, palins))
	}
}
