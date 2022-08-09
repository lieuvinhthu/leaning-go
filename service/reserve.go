package service

import (
	"context"
	"go-project/model"
	"go-project/repository"
)

type (
	ReserveService struct {
		repo *repository.ReserveRepo
	}

	ReserveServiceImpl interface {
		GetAllTable(ctx context.Context, date string) ([]*model.Reserve, error)
		CancelBooking(reserves []*model.Reserve, reqPhoneNumber string) []*model.Reserve
		UpdateBooking(reqReserve model.Reserve, reserves []*model.Reserve) []*model.Reserve
		CreateNewBooking(ctx context.Context, reserve model.Reserve) error
	}
)

func NewService(repo *repository.ReserveRepo) *ReserveService {
	return &ReserveService{
		repo: repo,
	}
}

func (s *ReserveService) GetAllTable(ctx context.Context, date string) ([]*model.Reserve, error) {
	filterReserves, err := s.repo.GetAllTable(ctx, date)
	if err != nil {
		return nil, err
	}
	return filterReserves, nil
}

func (s *ReserveService) CreateNewBooking(ctx context.Context, reserve model.Reserve) error {
	err := s.repo.CreateNewBooking(ctx, reserve)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReserveService) CancelBooking(ctx context.Context, reqPhoneNumber string) error {
	err := s.repo.CancelBooking(ctx, reqPhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReserveService) UpdateBooking(ctx context.Context, reserve model.Reserve) error {
	err := s.repo.UpdateBooking(ctx, reserve)
	if err != nil {
		return err
	}
	return nil
}
