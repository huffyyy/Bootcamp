package response

import "time"

// Response adalah wrapper standar untuk semua response API
type Response[T any] struct {
	Success   bool      `json:"success"`           // true jika berhasil, false jika gagal
	Message   string    `json:"message,omitempty"` // pesan sukses atauerror description
	Data      *T        `json:"data,omitempty"`    // payload utama (bisa nil jika error)
	Meta      *Meta     `json:"meta,omitempty"`    // pagination, count, dll(opsional)
	Errors    []Error   `json:"errors,omitempty"`  // detail error (untuk validation/multi-error)
	Timestamp time.Time `json:"timestamp"`         // waktu response
}

type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"` // "eq", "neq", "gt", "gte", "lt", "lte", "contains", "in", dll
	Value    interface{} `json:"value"`
}
type Meta struct {
	CurrentPage int      `json:"current_page,omitempty"`
	PageSize    int      `json:"page_size,omitempty"`
	TotalPages  int      `json:"total_pages,omitempty"`
	TotalItems  int      `json:"total_items,omitempty"`
	SortBy      string   `json:"sort_by,omitempty"`
	SortDir     string   `json:"sort_direction,omitempty"` // "asc" / "desc"
	Filters     []Filter `json:"filters,omitempty"`
}

// Error detail (untuk validation atau business error)
type Error struct {
	Field   string `json:"field,omitempty"` // nama field yang error (untuk validation)
	Code    string `json:"code,omitempty"`  // error code (misal: VALIDATION_ERROR, NOT_FOUND)
	Message string `json:"message"`         // deskripsi error
}
