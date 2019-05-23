package wxa

// Map returns a function which will run λ on an arbitrary list of interfaces
// and return a slice containing the results.
func Map(λ func(interface{}) (interface{}, error)) func(...interface{}) ([]interface{}, error) {
	return func(args ...interface{}) ([]interface{}, error) {
		var err error
		results := make([]interface{}, len(args))
		for i, arg := range args {
			if results[i], err = λ(arg); err != nil {
				return nil, err
			}
		}
		return results, nil
	}
}
