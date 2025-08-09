package util

import (
	"strings"

	"gorm.io/gorm"
)

type (
	PageFilter struct {
		PageSize    int  `json:"pageSize"`
		CurrentPage int  `json:"currentPage"`
		AllPages    bool `json:"allPages"`
	}

	SortFilter struct {
		SortField string `json:"sortField"`
		SortOrder string `json:"sortOrder"`
	}
)

func AddPagination(db *gorm.DB, pf PageFilter, sf SortFilter) *gorm.DB {
	if sf.SortField != "" {
		order := "asc"
		if strings.ToLower(sf.SortOrder) == "desc" {
			order = "desc"
		}
		db = db.Order(sf.SortField + " " + order)
	}

	if !pf.AllPages {
		if pf.PageSize <= 0 {
			pf.PageSize = 10
		}
		if pf.CurrentPage <= 0 {
			pf.CurrentPage = 1
		}
		offset := (pf.CurrentPage - 1) * pf.PageSize
		db = db.Offset(offset).Limit(pf.PageSize)
	}

	return db
}
