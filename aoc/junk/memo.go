package junk

func Memo[I comparable, O any](fn func(I) O) func(I) O {
	memo := make(map[I]O)
	return func(in I) O {
		res, found := memo[in]
		if !found {
			res = fn(in)
			memo[in] = res
		}
		return res
	}
}
