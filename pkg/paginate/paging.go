// data paging is a helper to
// create a pagination for data retrieve from repository layer
// Author : Robertus Adi Setyadharma - 9 September 2023
package paginate

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"gorm.io/gorm"
)

const (
	FirstPage          = 1
	PaginationMinLimit = 10
	SortAscending      = "asc"
	SortDescending     = "desc"
	DefaultSortColumn  = "id"
)

const (
	DEFAULT_SORT_COLUMN    = "id"
	DEFAULT_SORT_DIRECTION = "asc"
)

type DateFilter struct {
	StartDate time.Time
	EndDate   time.Time
}

type Datapaging struct {
	DateInTimestamp bool

	Limit int
	Page  int

	//OrderBy define the order property of the row, use it with format [field] [asc/desc].
	//Example []string{"distance desc"}
	OrderBy []string

	OrderByMulti []string

	//FilterColumn specify column as filter parameter
	FilterColumn string
	//FilterValue specify value of the column to be filtered
	FilterValue string

	DateLatest   *time.Time
	DateEarliest *time.Time

	DateBetweenPrefix string
}

// New will return a new pagination object specified with pagination,
// limit and order by field
func New(limit, page int, orderBy []string) Datapaging {
	return Datapaging{
		Limit:   limit,
		Page:    page,
		OrderBy: orderBy,
	}
}

// NoPagination return empty pagination
func NoPagination() Datapaging {
	return Datapaging{}
}

// IsNil check if the pagination object is empty
func (pagination *Datapaging) IsNil() bool {
	if !pagination.WithLimit() && !pagination.WithPageOffset() && !pagination.WithOrderBy() {
		return true
	}
	return false
}

func (pagination *Datapaging) GetOffset() int {
	return (pagination.Page - 1) * pagination.Limit
}

func (pagination *Datapaging) WithLimit() bool {
	if pagination.Limit != 0 {
		return true
	}
	return false
}

func (pagination *Datapaging) WithPageOffset() bool {
	if pagination.Page != 0 {
		return true
	}
	return false
}

func (pagination *Datapaging) WithOrderBy() bool {
	if len(pagination.OrderBy) > 0 {
		return true
	}
	return false
}

func (pagination *Datapaging) WithOrderByMulti() bool {
	if len(pagination.OrderByMulti) > 0 {
		return true
	}
	return false
}

func (pagination Datapaging) Between(earliestTime, latestTime *time.Time) Datapaging {
	pagination.DateEarliest = earliestTime
	pagination.DateLatest = latestTime
	return pagination
}

func (pagination *Datapaging) WithDateBetween() bool {
	if pagination.DateEarliest != nil && pagination.DateLatest != nil && !pagination.DateInTimestamp {
		return true
	}

	return false
}

func (pagination *Datapaging) WithDateTimeBetween() bool {
	if pagination.DateEarliest != nil && pagination.DateLatest != nil && pagination.DateInTimestamp {
		return true
	}

	return false
}

// BuildQueryGORM build datapaging for GORM DB instance
func (pagination *Datapaging) BuildQueryGORM(db *gorm.DB) *gorm.DB {

	if pagination.WithLimit() {
		db = db.Limit(pagination.Limit)
	}

	if pagination.WithPageOffset() {
		db = db.Offset(pagination.GetOffset())
	}

	if pagination.WithOrderBy() {
		db = db.Order(fmt.Sprintf("%s %s", pagination.OrderBy[0], pagination.OrderBy[1]))
	}

	if pagination.WithOrderByMulti() {
		for _, orderData := range pagination.OrderByMulti {
			db = db.Order(orderData)
		}
	}

	if pagination.WithDateBetween() {
		createdAtCol := "created_at"
		if pagination.DateBetweenPrefix != "" {
			createdAtCol = pagination.DateBetweenPrefix + "." + createdAtCol
		}

		db = db.Where(createdAtCol+" >= ? AND "+createdAtCol+" < ?", pagination.DateEarliest.Unix(),
			pagination.DateLatest.AddDate(0, 0, 1).Unix())
	}

	if pagination.WithDateTimeBetween() {
		createdAtCol := "created_at"
		if pagination.DateBetweenPrefix != "" {
			createdAtCol = pagination.DateBetweenPrefix + "." + createdAtCol
		}

		db = db.Where(createdAtCol+" >= ? AND "+createdAtCol+" < ?", pagination.DateEarliest,
			pagination.DateLatest.AddDate(0, 0, 1))
	}

	return db
}

func (pagination *Datapaging) BuildQueryGORMWithCustomDateColumn(db *gorm.DB, column string) *gorm.DB {

	if pagination.WithOrderBy() {
		db = db.Order(fmt.Sprintf("%s %s", pagination.OrderBy[0], pagination.OrderBy[1]))
	}

	if pagination.WithDateBetween() {
		db = db.Where(column+" >= ? AND "+column+" < ?", pagination.DateEarliest.Unix(), pagination.DateLatest.AddDate(0, 0, 1).Unix())
	}

	if pagination.WithDateTimeBetween() {
		db = db.Where(column+" >= ? AND "+column+" < ?", pagination.DateEarliest, pagination.DateLatest.AddDate(0, 0, 1))
	}

	if pagination.WithLimit() {
		db = db.Limit(pagination.Limit)
	}

	if pagination.WithPageOffset() {
		db = db.Offset(pagination.GetOffset())
	}

	return db
}

// BuildQuery will add the pagination syntax into the raw sqlQuery
func (pagination *Datapaging) BuildQuery(sqlQuery string) string {

	if pagination.WithOrderBy() {
		sqlQuery = sqlQuery + " ORDER BY"
		for i, order := range pagination.OrderBy {
			if len(pagination.OrderBy)-i == 1 {
				sqlQuery = sqlQuery + " " + order
			} else {
				sqlQuery = sqlQuery + " " + order + ","
			}
		}
	}

	if pagination.WithLimit() {
		sqlQuery = sqlQuery + " LIMIT " + strconv.Itoa(pagination.Limit)
	}

	if pagination.WithPageOffset() {
		pageOffset := (pagination.Limit * (pagination.Page)) - pagination.Limit
		sqlQuery = sqlQuery + " OFFSET " + strconv.Itoa(pageOffset)
	}

	return sqlQuery
}

type DataPagingResponseDTO struct {
	PageNumber       int         `json:"page_number"`
	PageSize         int         `json:"page_size"`
	Limit            int         `json:"limit"`
	TotalRecordCount int64       `json:"total_record_count"`
	Records          interface{} `json:"records"`
}

func (d *DataPagingResponseDTO) SetPageSize() *DataPagingResponseDTO {
	if d.Limit != 0 {
		d.PageSize = int(math.Ceil(float64(d.TotalRecordCount) / float64(d.Limit)))
	}
	return d
}
