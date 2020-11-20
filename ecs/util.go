package ecs

// with is a copy and breaks the cycle
func with(next func(fields ...Field), more ...func() Field) func(fields ...Field) {
	return func(fields ...Field) {
		tmp := make([]Field, 0, len(more))
		for _, f := range more {
			tmp = append(tmp, f())
		}

		tmp = append(tmp, fields...)
		next(tmp...)
	}
}

// WithName returns a new logger function which always prepends the name.
func WithName(next func(fields ...Field), name string) func(fields ...Field) {
	return with(next, func() Field {
		return Log(name)
	})
}

// WithTime returns a new logger function which always prepends the time.
func WithTime(next func(fields ...Field)) func(fields ...Field) {
	return with(next, Time)
}
