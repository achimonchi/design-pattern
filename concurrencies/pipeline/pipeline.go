package pipeline

func LaunchPipeline(amount int) int {
	first := generator(amount)
	second := power(first)
	third := sum(second)

	res := <-third
	return res
}

func generator(max int) <-chan int {
	outChan := make(chan int, 100)

	go func() {
		for i := 1; i <= max; i++ {
			outChan <- i
		}

		close(outChan)
	}()
	return outChan
}

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
