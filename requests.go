package rest

import (
	"strconv"
)

type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty"`
}

func ListOptionsParse(vars map[string]string, defaultPage int, defaultPerPage int) (int, int, error) {
	var page, perPage int
	if pageStr, ok := vars["page"]; ok {
		page64, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			return 0, 0, err
		}
		page = int(page64)
	} else {
		page = defaultPage
	}

	if perPageStr, ok := vars["per_page"]; ok {
		perPage64, err := strconv.ParseInt(perPageStr, 10, 64)
		if err != nil {
			return 0, 0, err
		}
		perPage = int(perPage64)
	} else {
		perPage = defaultPerPage
	}

	return page, perPage, nil
}
