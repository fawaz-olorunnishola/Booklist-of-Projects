package piscinego

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var count [10]int

	for n > 0 {
		digit := n % 10
		count[digit]++
		n /= 10
	}

	for i := 0; i <= 9; i++ {
		for count[i] > 0 {
			z01.PrintRune(rune(i + '0'))
			count[i]--
		}
	}
}
