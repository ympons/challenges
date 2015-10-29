package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	Fib []uint64 = []uint64{2, 8, 34, 144, 610, 2584, 10946, 46368, 196418, 832040, 3524578, 14930352, 63245986,
		267914296, 1134903170, 4807526976, 20365011074, 86267571272, 365435296162,
		1548008755920, 6557470319842, 27777890035288, 117669030460994, 498454011879264,
		2111485077978050, 8944394323791464, 37889062373143906}
	Sum []uint64 = []uint64{2, 10, 44, 188, 798, 3382, 14328, 60696, 257114, 1089154, 4613732, 19544084, 82790070,
		350704366, 1485607536, 6293134512, 26658145586, 112925716858, 478361013020,
		2026369768940, 8583840088782, 36361730124070, 154030760585064, 652484772464328,
		2763969850442378, 11708364174233842, 49597426547377748}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wout := bufio.NewWriter(os.Stdout)
	if scanner.Scan() {
		for t, _ := strconv.Atoi(scanner.Text()); t > 0 && scanner.Scan(); t-- {
			n, _ := strconv.ParseUint(scanner.Text(), 10, 64)
			var i int
			for i = 0; i < len(Fib) && Fib[i] < n; i++ {
			}
			wout.WriteString(strconv.FormatUint(Sum[i-1], 10))
			wout.WriteString("\n")
		}
		wout.Flush()
	}
}