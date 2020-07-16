package sqlorder

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type SqlOrder struct {
	Column string `json:"column"`
	Desc   bool   `json:"desc"`
	Sort   int    `json:"sort"`
}

type SqlOrderCond struct {
	Orders       []*SqlOrder
	registColumn map[string]string
}

func (oc *SqlOrderCond) RegistColumn(columnMap map[string]string) {
	oc.registColumn = columnMap
}

func (oc *SqlOrderCond) String() string {
	res := ""
	for _, item := range oc.Orders {
		columnName, found := oc.registColumn[item.Column]
		if !found {
			continue
		}

		if res != "" {
			res += ", "
		}
		if item.Desc {
			res += fmt.Sprintf("%s desc", columnName)
			continue
		}
		res += item.Column
	}
	return res
}

func GetSortCond(c *gin.Context) *SqlOrderCond {
	res := &SqlOrderCond{Orders: []*SqlOrder{}}
	sort, found := c.GetQuery("sort")
	if !found {
		return nil
	}

	for i, item := range strings.Split(sort, ",") {
		tmp := &SqlOrder{Sort: i}
		res.Orders = append(res.Orders, tmp)
		column := strings.Split(item, " ")
		if len(column) == 2 && column[1] == "desc" {
			tmp.Column = column[0]
			tmp.Desc = true
			continue
		}
		tmp.Column = column[0]
		tmp.Desc = false
	}

	return res
}
