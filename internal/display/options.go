package display

import "time"

type Options struct {
	Compact  bool
	JSON     bool
	Detailed bool
	Now      time.Time
}
