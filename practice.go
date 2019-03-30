package kenbunshoku

import (
	"fmt"
	"time"
)

func main() {
	createTime, _ := time.Parse(time.RFC3339, "2019-03-21T22:08:41+00:00")
	fmt.Println(createTime)
	if createTime.IsZero() {
		return
	}

	duration := time.Duration(24)
	fmt.Println(duration)
	from := createTime.Truncate(24 * time.Hour)
	fmt.Println("from:", from)
	to := createTime.Add(duration * time.Hour).Truncate(24 * time.Hour)
	fmt.Println("to:", to)
}
