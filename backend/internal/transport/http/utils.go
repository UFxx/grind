package http

import "math"

func GetPaginationResponse(totalElements int64, currentPage, limit int) PaginationResponse {

	totalPages := int(math.Ceil(float64(totalElements) / float64(limit)))
	nextPage := currentPage + 1

	if nextPage > totalPages {
		nextPage = 0
	}

	return PaginationResponse{
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		NextPage:    nextPage,
		Limit:       limit,
	}
}
