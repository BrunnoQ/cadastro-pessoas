package dto

// PaginationRequest represents pagination parameters
type PaginationRequest struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

// DefaultPagination returns default pagination values
func DefaultPagination() PaginationRequest {
	return PaginationRequest{
		Limit:  20,
		Offset: 0,
	}
}

// Validate validates and normalizes pagination parameters
func (p *PaginationRequest) Validate() {
	// Set defaults if not provided
	if p.Limit <= 0 {
		p.Limit = 20
	}
	if p.Offset < 0 {
		p.Offset = 0
	}

	// Enforce maximum limit to prevent abuse
	if p.Limit > 100 {
		p.Limit = 100
	}
}
