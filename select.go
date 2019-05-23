package wxa

// Select returns a function which takes some things and returns a slice
// containing those things for which λ returns true.
func Select(λ func(interface{}) (bool, error)) func(...interface{}) ([]interface{}, error) {
	return func(args ...interface{}) ([]interface{}, error) {
		r := []interface{}{}
		for _, arg := range args {
			pass, err := λ(arg)
			if err != nil {
				return r, err
			}
			if pass {
				r = append(r, arg)
			}
		}
		return r, nil
	}
}
