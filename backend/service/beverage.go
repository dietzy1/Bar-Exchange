package service

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type Beverage struct {
	Id               string       `json:"id" bson:"id"`
	Price            string       `json:"price" bson:"price"`
	BasePrice        float64      `json:"base_price" bson:"base_price"`
	Name             string       `json:"name" bson:"name"`
	PercentageChange int          `json:"percentage_change" bson:"percentage_change"`
	Type             BeverageType `json:"type" bson:"type"`
	Status           Status       `json:"status" bson:"status"`
}

type BeverageType int32

const (
	BEVERAGE_TYPE_UNSPECIFIED BeverageType = iota
	BEVERAGE_TYPE_BEER
	BEVERAGE_TYPE_COCKTAIL
	BEVERAGE_TYPE_SHOTS
)

type Status int32

const (
	STATUS_UNSPECIFIED Status = iota
	STATUS_INCREASING
	STATUS_DECREASING
	STATUS_NO_CHANGE
)

type beverageStore interface {
	GetBeverages(ctx context.Context) ([]Beverage, error)
	CreateBeverage(ctx context.Context, req Beverage) error
	UpdateBeverage(ctx context.Context, req Beverage) error
	DeleteBeverage(ctx context.Context, req Beverage) error
}

type beverageService struct {
	store beverageStore

	logger *zap.Logger
}

func NewBeverageService(store beverageStore, logger *zap.Logger) (*beverageService, error) {

	if store == nil {
		return nil, fmt.Errorf("BeverageStore is nil")
	}

	//Create beverage service
	beverageService := &beverageService{store: store, logger: logger}

	return beverageService, nil
}

func (s *beverageService) GetBeverages(ctx context.Context) ([]Beverage, error) {

	res, err := s.store.GetBeverages(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *beverageService) CreateBeverage(ctx context.Context, req Beverage) error {

	if err := ValidateCreateBeverage(req); err != nil {
		s.logger.Error("failed to validate beverage", zap.Error(err))
		return fmt.Errorf("failed to validate beverage: %w", err)
	}

	//Add missing fields
	req.Id = newID()
	req.PercentageChange = 0
	req.Status = STATUS_NO_CHANGE

	if err := s.store.CreateBeverage(ctx, req); err != nil {
		return fmt.Errorf("failed to create beverage: %w", err)
	}

	return nil

}

func ValidateCreateBeverage(input Beverage) error {

	var errorBuilder strings.Builder

	//Check for valid price
	if input.Price == "" {
		errorBuilder.WriteString("Invalid price\n")
	}

	//Check for valid name
	if input.Name == "" {
		errorBuilder.WriteString("Invalid name\n")
	}

	//Check for valid beverage type
	switch input.Type {
	case BEVERAGE_TYPE_BEER:
	case BEVERAGE_TYPE_COCKTAIL:
	case BEVERAGE_TYPE_SHOTS:
	default:
		errorBuilder.WriteString("Invalid beverage type\n")
	}

	if len(errorBuilder.String()) > 0 {
		return fmt.Errorf(errorBuilder.String())
	}

	return nil
}

func (s *beverageService) UpdateBeverage(ctx context.Context, req Beverage) error {

	if err := validateUpdateBeverage(req); err != nil {
		s.logger.Error("failed to validate beverage", zap.Error(err))
		return fmt.Errorf("failed to validate beverage: %w", err)
	}

	if err := s.store.UpdateBeverage(ctx, req); err != nil {
		return fmt.Errorf("failed to create beverage: %w", err)
	}

	return nil

}

func validateUpdateBeverage(input Beverage) error {

	var errorBuilder strings.Builder

	if input.Id == "" {
		errorBuilder.WriteString("Invalid id\n")
	}

	//Check for valid price
	if input.Price == "" {
		errorBuilder.WriteString("Invalid price\n")
	}

	//Check for valid name
	if input.Name == "" {
		errorBuilder.WriteString("Invalid name\n")
	}

	//Check for valid beverage type
	switch input.Type {
	case BEVERAGE_TYPE_BEER:
	case BEVERAGE_TYPE_COCKTAIL:
	case BEVERAGE_TYPE_SHOTS:
	default:
		errorBuilder.WriteString("Invalid beverage type\n")
	}

	//Check for valid status
	switch input.Status {
	case STATUS_INCREASING:
	case STATUS_DECREASING:
	case STATUS_NO_CHANGE:
	default:
		errorBuilder.WriteString("Invalid status\n")
	}

	if len(errorBuilder.String()) > 0 {
		return fmt.Errorf(errorBuilder.String())
	}

	return nil

}

func (s *beverageService) DeleteBeverage(ctx context.Context, req Beverage) error {

	if req.Id == "" {
		s.logger.Error("invalid id")
		return fmt.Errorf("invalid id")
	}

	if err := s.store.DeleteBeverage(ctx, req); err != nil {
		return fmt.Errorf("failed to delete beverage: %w", err)
	}

	return nil
}
