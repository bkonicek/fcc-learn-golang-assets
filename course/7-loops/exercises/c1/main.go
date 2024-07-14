package main

/*
Since we know that divisors and dividends will always mirror each other (e.g. 10/2 == 5 and 10/5 ==2)
we know that the dividend of the smallest divisor (meeting the conditions) divided into the message length will
also be the largest valid divisor.
*/
func getPacketSize(message string) int {
	// Since we're flipping the divisor (number of packets) and dividends (packet size), we can
	// skip 1 through 3 as they are either 1 or prime and we know
	// those aren't valid packet sizes.
	for i := 4; i < len(message); i++ {
		if isPrime(i) {
			continue
		}
		if len(message)%i == 0 {
			return len(message) / i
		}
	}
	return 0
}

func isPrime(num int) bool {
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// You shall not pass!! 20
// May the odds be ever in your favor!! 36
