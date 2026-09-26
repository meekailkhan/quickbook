package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/meekailkhan/quick-book/db"
)

const defaultLimit = 20

// GetAllInventoriesParams holds filter and pagination options for the service.
type GetAllInventoriesParams struct {
	Category string // optional — empty means no category filter
	Page     int    // 1-indexed; defaults to 1
}

// InventoryResult wraps the paginated response with metadata.
type InventoryResult struct {
	Inventories []db.ListInventoriesRow `json:"inventories"`
	Total       int64                   `json:"total"`
	Page        int                     `json:"page"`
	Limit       int                     `json:"limit"`
	TotalPages  int                     `json:"total_pages"`
}

// InventoryService handles inventory-related business logic.
type InventoryService struct {
	queries *db.Queries
}

// NewInventoryService creates a new InventoryService backed by the given *sql.DB.
func NewInventoryService(sqlDB *sql.DB) *InventoryService {
	return &InventoryService{
		queries: db.New(sqlDB),
	}
}

// GetAllInventories returns a paginated list of available inventories (20 per page).
// Filters by category when params.Category is non-empty.
func (s *InventoryService) GetAllInventories(ctx context.Context, params GetAllInventoriesParams) (*InventoryResult, error) {
	if params.Page < 1 {
		params.Page = 1
	}

	offset := int32((params.Page - 1) * defaultLimit)
	limit := int32(defaultLimit)

	var (
		inventories []db.ListInventoriesRow
		total       int64
		err         error
	)

	inventories, err = s.queries.ListInventories(ctx, db.ListInventoriesParams{
		Category: params.Category,
		Limit:    limit,
		Offset:   offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list inventories: %w", err)
	}

	totalPages := int(total) / defaultLimit
	if int(total)%defaultLimit != 0 {
		totalPages++
	}

	return &InventoryResult{
		Inventories: inventories,
		Total:       total,
		Page:        params.Page,
		Limit:       defaultLimit,
		TotalPages:  totalPages,
	}, nil
}
