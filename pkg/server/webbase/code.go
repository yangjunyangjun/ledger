package webbase

const Success = 20000

const (
	InvalidRequest = 50000+iota
	JsonUnmarshal
	InvalidToken
	DB
)
