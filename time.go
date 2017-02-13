package goutils

import (
	"fmt"
	"time"
)

var (

	// TimeIntervalWeekly 时间区间 周
	TimeIntervalWeekly = 1

	// TimeIntervalDaily  时间区间 日
	TimeIntervalDaily = 2
)

// TimeIntervalBlock 时间区间
type TimeIntervalBlock struct {
	Begin time.Time
	End   time.Time
}

// TimeIntervalBlocks 时间区间集合
type TimeIntervalBlocks struct {
	Model  string
	Blocks []*TimeIntervalBlock
}

// GetTimeNowRFC3339 获取RFC3339的现在格式
func GetTimeNowRFC3339() string {
	return time.Now().Format(time.RFC3339)
}

// GetTimeIntervalBlocks returns the datetime blocks of special model(weekly,daily) between begin and end params
// model {1:weekly , 2: daily}
func GetTimeIntervalBlocks(model int, begin, end time.Time) (blocks *TimeIntervalBlocks, err error) {

	if !IntIsIn(model, TimeIntervalDaily) {
		return nil, fmt.Errorf("only support model, daily:%d, unsupported model:%d", TimeIntervalDaily, model)
	}

	if end.After(begin) {

		blocks = &TimeIntervalBlocks{}
		theDayAfterLastDayBegin := GetOneDayBeginOfTime(end).Add(24 * time.Hour)
		fmt.Println(theDayAfterLastDayBegin)
		for temp := begin; temp.Before(theDayAfterLastDayBegin); temp = temp.Add(24 * time.Hour) {
			fmt.Println(GetOneDayBeginOfTime(temp))
			fmt.Println(GetOneDayEndOfTime(temp))
			blocks.Blocks = append(blocks.Blocks, &TimeIntervalBlock{GetOneDayBeginOfTime(temp), GetOneDayEndOfTime(temp)})
		}
		return blocks, nil

	}

	return nil, fmt.Errorf("bad begin and end time, start must before end ")

}

// GetOneDayBeginOfTime returns the begin of the t
func GetOneDayBeginOfTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetOneDayEndOfTime returns the end of the t
func GetOneDayEndOfTime(t time.Time) time.Time {
	return GetOneDayBeginOfTime(t).Add(24 * time.Hour).Add(-1 * time.Nanosecond)
}
