package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"

	t "github.com/cznic/mathutil"
)

func Solve(a []string) int {
	i := new(big.Int)
	n := new(big.Int)
	plusone := new(big.Int)
	fmt.Sscan("1", plusone)
	cases := 0
	fmt.Sscan(a[0], i)
	fmt.Sscan(a[1], n)
	for _, palindrome := range GeneratePalindromes(i, n) {
		if isBigSquare(&palindrome) {
			root := t.SqrtBig(&palindrome)
			if isBigPalindrome(root) {
				cases += 1
			}
		}
	}
	return cases
}

func GeneratePalindromesWithLength(l int) []*big.Int {
	one := big.NewInt(1)
	var palindromes []*big.Int
	if l == 1 {
		for i := 1; i < 10; i++ {
			x := big.NewInt(int64(i))
			palindromes = append(palindromes, x)
		}
		return palindromes
	}
	if l%2 == 1 {
		half_length := (l - 1) / 2
		for x := big.NewInt(int64(0)); x.Cmp(big.NewInt(10)) < 0; x = x.Add(x, one) {
			for y := big.NewInt(int64(math.Pow(10, float64(half_length-1)))); y.Cmp(big.NewInt(int64(math.Pow(10, float64(half_length))))) < 0; y = y.Add(y, one) {
				palindrome := new(big.Int)
				palstring := y.String() + x.String() + Reverse(y.String())
				fmt.Sscan(palstring, palindrome)
				palindromes = append(palindromes, palindrome)
			}
		}
	} else {
		half_length := l / 2
		for x := big.NewInt(int64(math.Pow(10, float64(half_length-1)))); x.Cmp(big.NewInt(int64(math.Pow(10, float64(half_length))))) < 0; x = x.Add(x, one) {
			palindrome := new(big.Int)
			palstring := x.String() + Reverse(x.String())
			fmt.Sscan(palstring, palindrome)
			palindromes = append(palindromes, palindrome)
		}
	}
	return palindromes
}

func GeneratePalindromes(min *big.Int, max *big.Int) []big.Int {
	min_len := len(min.String())
	max_len := len(max.String())
	var palindromes []big.Int
	for l := min_len; l <= max_len; l++ {
		for _, x := range GeneratePalindromesWithLength(l) {
			if x.Cmp(min) >= 0 && x.Cmp(max) <= 0 {
				palindromes = append(palindromes, *x)
			}
		}
	}
	return palindromes
}

func isBigPalindrome(s *big.Int) bool {
	i := new(big.Int)
	fmt.Sscan(Reverse(s.String()), i)
	return s.Cmp(i) == 0
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

func isBigSquare(s *big.Int) bool {
	n := new(big.Int)
	root := t.SqrtBig(s)
	return s.Cmp(n.Mul(root, root)) == 0
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

func main() {
	argsWithoutProgram := os.Args[1]
	lines := GrabLines(argsWithoutProgram)
	cases, _ := strconv.Atoi(lines[0])
	for i := 1; i <= cases; i++ {
		line := strings.Split(lines[i], " ")
		fmt.Printf("Case #%d: %d\n", i, Solve(line))
	}
}