package logger

// type 'Field' represents a logging field
type Field struct {
	Key   string
	Value any
}

// function 'String' creates a string field
func String(key, val string) Field {
	return Field{Key: key, Value: val}
}

// function 'Int' creates an integer field
func Int(key string, val int) Field {
	return Field{Key: key, Value: val}
}

// function 'Bool' creates a boolean field
func Bool(key string, val bool) Field {
	return Field{Key: key, Value: val}
}

// function 'Map' creates a map field
func Map(key string, val map[string]any) Field {
	return Field{Key: key, Value: val}
}
