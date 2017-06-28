package utils

import "fmt"

// ByteToBinaryString Convert byte to binary string 8 digits
// Ex. 1 => 00000001
func ByteToBinaryString(byte byte) string {
	return fmt.Sprintf("%08b", byte)
}
