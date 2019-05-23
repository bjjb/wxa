package wxa

// A MapFunc is a function suitable for a Map invokation.
type MapFunc func(interface{}) (interface{}, error)

// A MustMapFunc is a function suitable for a MustMap invokation; rather than
// returning an error, it'll panic.
type MustMapFunc func(interface{}) interface{}

// Map takes a map function and an arbitrary number of arguments, and returns
// a slice of interfaces (which has the same length as the var-args) and an
// error which represents the first failed mapFunc application.
func Map(f MapFunc, args ...interface{}) ([]interface{}, error) {
	var err error
	results := make([]interface{}, len(args))
	for i, arg := range args {
		if results[i], err = f(arg); err != nil {
			return nil, err
		}
	}
	return results, nil
}

// MustMap takes a map function and an arbitrary number of arguments and
// applies the function to them, returning a slice of results.
func MustMap(f MustMapFunc, args ...interface{}) []interface{} {
	results := make([]interface{}, len(args))
	for i, arg := range args {
		results[i] = f(arg)
	}
	return results
}
