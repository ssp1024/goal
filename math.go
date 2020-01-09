package goal

func MaxInt(i int, args ...int) int {
	rv := i
	for _, n := range args {
		if n > rv {
			rv = n
		}
	}

	return rv
}
