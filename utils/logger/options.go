package logger

type Opts struct {
	field map[string]interface{}
}

type Options = func(*Opts)

func Field(key string, value interface{}) Opts {
	return Opts{
		field: map[string]interface{}{key: value},
	}
}
