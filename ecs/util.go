package ecs

// with is a copy and breaks the cycle
func with(next func(fields ...interface{}), more ...func() interface{}) func(fields ...interface{}) {
	return func(fields ...interface{}) {
		tmp := make([]interface{}, 0, len(more))
		for _, f := range more {
			tmp = append(tmp, f())
		}

		tmp = append(tmp, fields...)
		next(tmp...)
	}
}

// WithName returns a new logger function which always prepends the name.
func WithName(next func(fields ...interface{}), name string) func(fields ...interface{}) {
	return with(next, func() interface{} {
		return Log(name)
	})
}

// WithTime returns a new logger function which always prepends the time.
func WithTime(next func(fields ...interface{})) func(fields ...interface{}) {
	return with(next, func() interface{} {
		return Time
	})
}
