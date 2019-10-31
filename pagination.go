package nfeio

type pagination struct {
	TotalResults int64 `json:"totalResults"`
	TotalPages   int64 `json:"totalPages"`
	Page         int64 `json:"page"`
}
