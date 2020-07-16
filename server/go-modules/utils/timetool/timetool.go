package timetool

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	TimeDefaultLayout    string = "2006-01-02 15:04:05"
	TimeFormatYYMMDDhhmm string = "2006-01-02 15:04"
	TimeFormatYYMMDD     string = "2006-01-02"
)

func GetNowStringByUTC() string {
	return time.Now().UTC().Format(TimeDefaultLayout)
}

func GetNowByUTC() time.Time {
	return time.Now().UTC()
}

func GetDateByUTC() time.Time {
	now := time.Now().UTC().Format(TimeFormatYYMMDD)
	date, _ := time.Parse(TimeFormatYYMMDD, now)
	return date
}

type CreatedTimeCond struct {
	StartTime *time.Time `json:"start_time" db_column:"created_at" db_op:">="`
	EndTime   *time.Time `json:"end_time"   db_column:"created_at" db_op:"<"`
}

func (ctc *CreatedTimeCond) GetFromContext(c *gin.Context) error {
	if sTime, exist := c.GetQuery("start_time"); exist {
		startTime, err := time.Parse(TimeDefaultLayout, sTime)
		if err != nil {
			return err
		}
		ctc.StartTime = &startTime
	}

	if eTime, exist := c.GetQuery("end_time"); exist {
		endTime, err := time.Parse(TimeDefaultLayout, eTime)
		if err != nil {
			return err
		}
		ctc.EndTime = &endTime
	}

	if err := ctc.CheckTime(); err != nil {
		return err
	}
	return nil
}

func (ctc *CreatedTimeCond) CheckTime() error {
	if ctc.StartTime == nil && ctc.EndTime == nil {
		return nil
	}
	if (ctc.StartTime != nil && ctc.EndTime == nil) ||
		(ctc.StartTime == nil && ctc.EndTime != nil) {
		return fmt.Errorf("Error: start time or end time is null.")
	}
	if ctc.StartTime.After(*ctc.EndTime) {
		return fmt.Errorf("Error: start time after end time.")
	}
	return nil
}

type PayoutTimeCond struct {
	PayoutStartTime *time.Time `json:"payout_start_time" db_column:"payout_at" db_op:">="`
	PayoutEndTime   *time.Time `json:"payout_end_time"   db_column:"payout_at" db_op:"<"`
}

func (ptc *PayoutTimeCond) GetFromContext(c *gin.Context) error {
	if sTime, exist := c.GetQuery("payout_start_time"); exist {
		startTime, err := time.Parse(TimeDefaultLayout, sTime)
		if err != nil {
			return err
		}
		ptc.PayoutStartTime = &startTime
	}

	if eTime, exist := c.GetQuery("epayout_nd_time"); exist {
		endTime, err := time.Parse(TimeDefaultLayout, eTime)
		if err != nil {
			return err
		}
		ptc.PayoutEndTime = &endTime
	}

	if err := ptc.CheckTime(); err != nil {
		return err
	}
	return nil
}

func (ptc *PayoutTimeCond) CheckTime() error {
	if ptc.PayoutStartTime == nil && ptc.PayoutEndTime == nil {
		return nil
	}
	if (ptc.PayoutStartTime != nil && ptc.PayoutEndTime == nil) ||
		(ptc.PayoutStartTime == nil && ptc.PayoutEndTime != nil) {
		return fmt.Errorf("Error: start time or end time is null.")
	}
	if ptc.PayoutStartTime.After(*ptc.PayoutEndTime) {
		return fmt.Errorf("Error: start time after end time.")
	}
	return nil
}

type Timestamp time.Time

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(*t).Unix(), 10)), nil
}

func (t *Timestamp) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return
}

func (t *Timestamp) String() string { return time.Time(*t).Format(TimeDefaultLayout) }

func (t *Timestamp) Time() time.Time { return time.Time(*t).UTC() }
