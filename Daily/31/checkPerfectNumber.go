package Daily31

// CheckPerfectNumber checks if a number is a perfect number in O(√(log(n)))
func CheckPerfectNumber(num int) bool {
	return checkPerfectNumberSqrtLog(num)
}
func checkPerfectNumberSqrt(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return sum == num
}

// checkPerfectNumberSqrtLog checks if a number is a perfect number in O(√(log(n))) by counting bits
// Description of perfect numbers: https://www.personal.psu.edu/sxt104/class/Math140H/PerfectNum.html
func checkPerfectNumberSqrtLog(num int) bool {
	// if num == 1 {
	// 	return false
	// } // O(1) until here

	// count total and set bits
	totalBits, setBits, bits := 0, 0, num
	for bits > 0 {
		totalBits++
		if bits&1 == 1 {
			setBits++
		}
		bits >>= 1
	} // O(log n) until here

	// check if setBits is unsetBits-1
	if setBits != totalBits-setBits+1 {
		return false
	}
	return isPrime(setBits) // O(√(log(n)))
}
