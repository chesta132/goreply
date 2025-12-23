package reply

import (
	"reflect"
	"strings"
)

// Create pagination information in meta data with total based.
//
// Example:
//
//	rp.Success(datas).PaginateTotal(limit, current, total).OkJSON()
func (r *Reply) PaginateTotal(limit, current, total int) *Reply {
	if limit <= 0 {
		limit = 1
	}

	t := strings.ToLower(string(r.c.paginationType))
	var hasNext bool
	var next int

	if t == "page" {
		totalPages := (total + limit - 1) / limit
		hasNext = current < totalPages
		if hasNext {
			next = current + 1
		}
	} else {
		hasNext = current+limit < total
		if hasNext {
			next = current + limit
		}
	}

	r.m.Meta.Pagination = &Pagination{
		Next:    next,
		HasNext: hasNext,
		Current: current,
		Total:   total,
	}
	return r
}

// Create pagination information in meta data with cursor based (data to send + 1).
// Auto cut data.
//
// Example:
//
//	rp.Success(datas).PaginateCursor(limit, current).OkJSON()
func (r *Reply) PaginateCursor(limit, current int) *Reply {
	if limit <= 0 {
		limit = 1
	}

	t := strings.ToLower(string(r.c.paginationType))

	v := reflect.ValueOf(r.m.Data)
	if v.Kind() != reflect.Slice {
		return r
	}
	dataLen := v.Len()
	hasNext := dataLen > limit

	if hasNext {
		r.m.Data = v.Slice(0, limit).Interface()
	}

	var next int
	if hasNext {
		if t == "page" {
			next = current + 1
		} else {
			next = current + limit
		}
	}

	r.m.Meta.Pagination = &Pagination{
		Next:    next,
		HasNext: hasNext,
		Current: current,
	}
	return r
}
