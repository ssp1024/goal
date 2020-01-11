package goal

//MaxInt return max integer value of all argments.
func MaxInt(i int, args ...int) int {
	rv := i
	for _, n := range args {
		if n > rv {
			rv = n
		}
	}

	return rv
}
