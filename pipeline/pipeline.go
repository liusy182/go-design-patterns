package pipeline

func LaunchPipeline(amount int) int {
	return <-sum(power(generator(amount)))
}

// no 1
func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
		close(outChInt)
	}()
	return outChInt
}

// no 2
func power(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

// no 3
func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
