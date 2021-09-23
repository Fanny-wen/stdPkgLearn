package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//timeNowDemo()
	//timestampDemo()
	//timestampToTimeDemo(time.Now().Unix())
	//operationTimeDemo()
	//timeFormatDemo()
	//parseStringToTimeDemo()
	timeRoundTruncateDemo()
}

func timeNowDemo() {
	// 获取当前时间
	fmt.Println(os.Executable())
	now := time.Now()
	fmt.Printf("time.Now: %v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	nanosecond := now.Nanosecond()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d-%09d\n", year, month, day, hour, minute, second, nanosecond)
	fmt.Println("======================================================")
}

func timestampDemo() {
	// 时间戳
	now := time.Now()
	timestamp1 := now.Unix()
	now.IsZero()
	timestamp2 := now.UnixNano()
	fmt.Printf("current timestamp: %v\n", timestamp1)
	fmt.Printf("current nano timestamp: %v\n", timestamp2)
	fmt.Println("======================================================")
}

func timestampToTimeDemo(timestamp int64) {
	// 时间戳转为时间格式
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)                 // 2021-08-18 10:37:52 +0800 CST
	fmt.Printf("timeObj: %T\n", timeObj) // timeObj: time.Time
	fmt.Println("======================================================")
}

func operationTimeDemo() {
	now := time.Now()
	// 时间操作
	// Add  时间 + 时间
	date := time.Date(2010, 8, 18, 10, 9, 10, 0, time.Local)
	fmt.Println(date)
	later := date.Add(time.Hour * 24 * 365 * 10)
	fmt.Printf("%v 24Hour later: %v\n", date, later)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	// Sub 两个时间之差
	before := now.Sub(later)
	fmt.Printf("now - date: %v\n", before)
	fmt.Println("------------------------------------------------------")

	// Equal 判断两个时间是否相同, 会考虑时区的影响
	ret := now.Equal(date)
	fmt.Printf("quesion: now == date? answer: %t, %v\n", ret, ret)
	fmt.Println("======================================================")

	// Before 如果t代表的时间点在u之前，返回真；否则返回假。
	ret = now.Before(date)
	fmt.Printf("quesion: now < date? answer: %t, %v\n", ret, ret)

	// After 如果t代表的时间点在u之后，返回真；否则返回假。
	ret = now.After(date)
	fmt.Printf("quesion: now > date? answer: %t, %v\n", ret, ret)
}

func timeFormatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func parseStringToTimeDemo() {
	// Parse
	// time.Now() 的时区是 time.Local，而 time.Parse 解析出来的时区却是 time.UTC（可以通过 Time.Location() 函数知道是哪个时区）。在中国，它们相差 8 小时。
	timeObj, err := time.Parse("2006-01-02 15:04", "2012-12-12 04:40")
	if err != nil {
		fmt.Printf("time.Parse failed, err: %v\n", err)
		return
	}
	fmt.Printf("time.Parse success, value: %v, 时区: %q\n", timeObj, timeObj.Location())
	fmt.Println("======================================================")

	// ParseInLocation
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("time.LoadLocation failed, err: %v\n", err)
		return
	}
	timeObj, err = time.ParseInLocation("2006-01-02 15:04", "2016-12-22 11:04", loc)
	if err != nil {
		fmt.Printf("time.ParseInLocation failed, err: %v\n", err)
		return
	}
	fmt.Printf("time.ParseInLocation success, value: %v, 时区: %q\n", timeObj, timeObj.Location())
}

func timeRoundTruncateDemo() {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-09-23 22:25:35", time.Local)
	fmt.Println(t)
	// 整点（向下取整）
	fmt.Println(t.Truncate(1 * time.Hour))
	// 整点（最接近）
	fmt.Println(t.Round(1 * time.Hour))
	// 整分（向下取整）
	fmt.Println(t.Truncate(1 * time.Minute))
	// 整分（最接近）
	fmt.Println(t.Round(1 * time.Minute))

	fmt.Println("==============================")
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:00:00"), time.Local)
	fmt.Println(t2)
}
