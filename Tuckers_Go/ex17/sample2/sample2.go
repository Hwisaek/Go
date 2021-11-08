package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	const longForm = "Jan 2, 2006 at 3:04 pm"
	t1, _ := time.ParseInLocation(longForm, "Jun 13, 2021 at 10:00 pm", loc)
	fmt.Println(t1, "\t", t1.Location(), "\t", t1.UTC())

	const shortForm = "2006-Jan-02"
	t2, _ := time.Parse(shortForm, "2021-Jun-14") // 그리니치 천문대 시각 기준. UTC 시각 기준
	fmt.Println(t2, "\t", t2.Location())

	t3, err := time.Parse("2021-06-01 15:20:21", "2021-06-14 20:04:05")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3, "\t", t3.Location())

	d := t2.Sub(t1)
	fmt.Println(d)
}
