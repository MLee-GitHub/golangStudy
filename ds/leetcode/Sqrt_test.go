package leetcode

import (
	"fmt"
	"testing"
)

//¡™£¢∞§¶•ªº–≠œ∑´´†¥¨ˆˆπ“‘«åß∂ƒ©˙∆˚¬…æΩ≈ç√∫˜˜≤≥ç∫√çß
func TestNewtonSqrt(t *testing.T) {
	var x float64 = 123
	var precision int = 3
	//_, _ = fmt.Scan(&x, &precision)
	ans := NewtonSqrt(x, precision)
	fmt.Printf("√%f=%f\n", x, ans)
}
