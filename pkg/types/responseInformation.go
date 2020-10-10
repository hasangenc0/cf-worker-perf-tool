package types

import "time"

type ResponseInfo struct {
	Size       int
	Time       time.Duration
	Succeed    bool
	StatusCode int
}
