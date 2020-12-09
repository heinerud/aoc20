package itertools

func Combinations(list []int, n int, c chan []int) {
	if n == 1 {
		for _, x := range list {
			c <- []int{x}
		}
	} else if len(list) == n {
		c <- list
	} else if len(list) > n {
		for i, x := range list {
			cc := make(chan []int)
			go Combinations(list[i+1:], n-1, cc)
			for y := range cc {
				sr := []int{x}
				sr = append(sr, y...)
				c <- sr
			}
		}
	}
	close(c)
}
