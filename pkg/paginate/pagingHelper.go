package paginate

import (
	"strconv"
	"time"

	"github.com/kurdilesmana/backend-product-service/internal/core/models/helperModel"
	"github.com/spf13/cast"
)

func prepareSortBy(param string, allowedSortColumns []string) string {

	sortByFound := false
	for _, allowedSortColumn := range allowedSortColumns {
		if allowedSortColumn == param {
			sortByFound = true
			break
		}
	}

	if !sortByFound {
		return DefaultSortColumn
	}

	return param
}

func prepareSortDirection(param string) string {

	allowedSortDirections := []string{SortAscending, SortDescending}

	sortDirectionFound := false
	for _, allowedSortDirection := range allowedSortDirections {
		if allowedSortDirection == param {
			sortDirectionFound = true
			break
		}
	}

	if !sortDirectionFound {
		return SortAscending
	}

	return param
}

func PreparePagination(params map[string]string, allowedSortColumns []string) (paging Datapaging) {

	search := params["search"]
	sortBy := params["sort_by"]
	sortDirection := params["sort_direction"]
	page := cast.ToInt(params["page"])

	// Set default value for limit
	limit := PaginationMinLimit

	// Check if limit is provided
	if limitStr, ok := params["limit"]; ok && limitStr != "" {
		l, err := strconv.Atoi(limitStr)

		// Check if limit is valid
		if err == nil && l > 0 {
			limit = l
		} else {
			limit = PaginationMinLimit
		}
	}

	if page < FirstPage {
		page = FirstPage
	}

	paging = Datapaging{
		Limit:       limit,
		Page:        page,
		OrderBy:     []string{prepareSortBy(sortBy, allowedSortColumns), prepareSortDirection(sortDirection)},
		FilterValue: search,
	}

	return paging
}

func preparePaginationAndFilter(params map[string]string, allowedSortColumns []string) (paging Datapaging, filter helperModel.BaseFilter) {

	search := params["search"]
	sortBy := params["sortBy"]
	sortDirection := params["sortDirection"]

	filter = helperModel.BaseFilter{
		SortBy:        prepareSortBy(sortBy, allowedSortColumns),
		SortDirection: prepareSortDirection(sortDirection),
		Search:        search,
	}

	limit := cast.ToInt(params["limit"])
	page := cast.ToInt(params["page"])
	startDate := params["startDate"]
	endDate := params["endDate"]

	if page < FirstPage {
		page = FirstPage
	}

	if limit < PaginationMinLimit {
		limit = PaginationMinLimit
	}

	if startDate == "" {
		startDate = time.Now().Add(-cast.ToDuration("24h") * 7).Format("2006-01-02")
	}

	if endDate == "" {
		endDate = time.Now().Format("2006-01-02")
	}

	startTime, _ := time.Parse("2006-01-02 15:04:05", startDate+" 00:00:00")
	endTime, _ := time.Parse("2006-01-02 15:04:05", endDate+" 23:59:00")

	paging = Datapaging{DateEarliest: &startTime, DateLatest: &endTime, Page: page, Limit: limit}

	return paging, filter
}
