package model

import "fmt"

const SizeOfReceipt = 52

const (
	Dash       = "----------------------------------------------------"
	DoubleDash = "===================================================="
	Space      = "                                                    "
)

func CenterlizeReceipt(s string) string {
	padding := (SizeOfReceipt - len(s)) / 2
	return fmt.Sprintf("%*s%s%*s\n", padding, "", s, padding, "")
}

func JustifyRecept(left string, right string) string {
	ans := Space[len(left):]
	ans = left + ans
	ans = ans[:len(ans)-len(right)]
	ans = ans + right
	return ans
}
