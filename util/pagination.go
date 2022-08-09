package util

import (
	"strconv"
	model "task-scheduler/datamodel"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) (m *model.Pagination, err error) {
	// set default pagination if not set
	pagination := &model.Pagination{
		Limit:   10,
		Page:    1,
		OrderBy: "id desc",
	}

	// get url query
	query := c.Request.URL.Query()

	for k, v := range query {
		queryValue := v[len(v)-1]
		switch k {
		case "limit":
			pagination.Limit, err = strconv.Atoi(queryValue)
			if err != nil {
				return nil, err
			}
		case "page":
			pagination.Page, err = strconv.Atoi(queryValue)
			if err != nil {
				return nil, err
			}
		case "order_by":
			pagination.OrderBy = queryValue
		}
	}

	return pagination, nil
}
