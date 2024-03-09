package main

func bulkSend(numMessages int) float64 {
	if numMessages <= 0 {
		return 0.0
	}

	sum := 0.0
	for i := 0; i < numMessages; i++ {
		sum += 1 + float64(i)*0.01
	}
	return sum
}
