package wxa

// Find returns a function which, when given arguments, returns the first
// argument for which λ returns true.
func Find(λ func(interface{}) (bool, error)) func(...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		for _, arg := range args {
			r, err := λ(arg)
			if err != nil {
				return nil, err
			}
			if r {
				return arg, nil
			}
		}
		return nil, nil
	}
}
