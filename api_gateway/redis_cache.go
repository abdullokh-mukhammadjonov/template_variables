package apigateway

import "errors"

var (
	ErrCacheMiss   = errors.New("rediscache.miss")
	ErrInvalidDate = errors.New("invalid date, date format `YYYY-MM-DD`")
)