package util

import "time"

func TimeToYMD(t *time.Time) string {
	if t == nil {
		return ""
	} else {
		return t.Format("20060102150405")
	}
}
