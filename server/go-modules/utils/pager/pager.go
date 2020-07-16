package pager

import (
	"math"

	"github.com/gin-gonic/gin"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/utils/converttool"
)

var (
	DefaultPageIndex int = 1
	DefaultPageSize  int = 25
)

type PagerCond struct {
	PageSize  int
	PageIndex int
}

func GetPagerCond(c *gin.Context) *PagerCond {
	pc := &PagerCond{}
	if pi, found := c.GetQuery("page_index"); !found {
		pc.PageIndex = DefaultPageIndex
	} else {
		pc.PageIndex = converttool.GetStringToInt(pi, DefaultPageIndex)
	}

	if ps, found := c.GetQuery("page_size"); !found {
		pc.PageSize = DefaultPageSize
	} else {
		pc.PageSize = converttool.GetStringToInt(ps, DefaultPageSize)
	}
	return pc
}

type Pager struct {
	TotalNums int         `json:"total_nums"`
	MaxPages  int         `json:"max_pages"`
	PageSize  int         `json:"page_size"`
	PageIndex int         `json:"page_index"`
	From      int         `json:"from"`
	To        int         `json:"to"`
	Data      interface{} `json:"data"`
}

func GetPager(pi, ps, totalCount int) *Pager {
	res := &Pager{
		TotalNums: totalCount,
		PageIndex: pi,
		PageSize:  ps,
		From:      1,
		To:        0,
		MaxPages:  0,
	}

	if totalCount != 0 {
		res.From = calFrom(pi, ps)
		res.To = calTo(pi, ps, totalCount)
		res.MaxPages = calPages(totalCount, ps)
	}
	return res
}

func calFrom(pi, ps int) int {
	return (pi-1)*ps + 1
}

func calTo(pi, ps, totalCount int) int {
	max := pi * ps
	if max > totalCount {
		return totalCount
	}
	return max
}

func calPages(totalCount, ps int) int {
	return int(math.Ceil(float64(totalCount) / float64(ps)))
}

type PagerCount struct {
	Total int
}
