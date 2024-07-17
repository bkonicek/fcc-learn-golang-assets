package main

/*
Since we know that divisors and dividends will always mirror each other (e.g. 10/2 == 5 and 10/5 ==2)
we know that the dividend of the smallest divisor (meeting the conditions) divided into the message length will
also be the largest valid divisor.
*/
func getPacketSize(message string) int {
	messageLength := len(message)

	// Start from the largest possible packet size and move downwards
	for packetSize := messageLength; packetSize > 0; packetSize-- {
		if messageLength%packetSize == 0 {
			// Calculate number of packets
			numPackets := messageLength / packetSize

			// Check if numPackets is composite (not 1 or prime)
			if numPackets > 1 && !isPrime(numPackets) {
				return packetSize
			}
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
