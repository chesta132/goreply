package reply

import (
	"reflect"
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

	var hasNext bool
	var next int

	// calc has next & next value by pagination type
	if r.c.PaginationType == PaginationPage {
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

	v := reflect.ValueOf(r.m.Data)

	// returns zero value prevents panic
	if !v.IsValid() {
		return r
	}

	// unwrap value
	for v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// guard type
	if v.Kind() != reflect.Slice {
		return r
	}

	dataLen := v.Len()
	hasNext := dataLen > limit

	// cut data
	if hasNext {
		r.m.Data = v.Slice(0, limit).Interface()
	}

	// set next by pagination type
	var next int
	if hasNext {
		if r.c.PaginationType == PaginationPage {
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
