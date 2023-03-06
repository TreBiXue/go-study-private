package request

import "time"

type LoginByIdRequest struct {
	ID         string    `json:"id"`
	AccessTime time.Time `json:"access_time"`
}

func (r *LoginByIdRequest) Validate() error {
	if r.AccessTime.IsZero() {
		r.AccessTime = time.Now()
	}
	return nil
}
