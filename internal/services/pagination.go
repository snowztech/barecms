package services

import "barecms/internal/models"

func pagination(total int64, page, limit int) models.Pagination {
	return models.Pagination{Page: page, Limit: limit, Total: total, TotalPages: int((total + int64(limit) - 1) / int64(limit))}
}
