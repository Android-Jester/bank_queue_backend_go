package transaction

import "time"

type Transaction struct {
	Detail      string
	ServerId    string
	NationalId  string
	Duration    float64
	CreatedDate time.Time
}
