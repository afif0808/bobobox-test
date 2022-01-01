package usecase

import (
	"context"

	"github.com/afif0808/bobobox_test/models"
	"github.com/afif0808/bobobox_test/payloads"
	"github.com/afif0808/bobobox_test/pkg/structs"
)

type roomPriceRepository interface {
	InsertRoomPrice(ctx context.Context, rp *models.RoomPrice) error
	InsertManyRoomPrice(ctx context.Context, rps []models.RoomPrice) error
	FetchRoomTypePrices(ctx context.Context, typeID int64) ([]models.RoomPrice, error)
	UpdateRoomPrice(ctx context.Context, rp models.RoomPrice, id int64) error
	DeleteRoomPrice(ctx context.Context, id int64) error
	GetRoomPrice(ctx context.Context, id int64) (models.RoomPrice, error)
}

type roomTypeRepository interface {
	GetRoomType(ctx context.Context, id int64) (models.RoomType, error)
}

type RoomPriceUsecase struct {
	rpRepo roomPriceRepository
	rtRepo roomTypeRepository
}

func NewRoomPriceUsecase(rpRepo roomPriceRepository, rtRepo roomTypeRepository) *RoomPriceUsecase {
	uc := RoomPriceUsecase{
		rpRepo: rpRepo,
		rtRepo: rtRepo,
	}
	return &uc
}

func (uc *RoomPriceUsecase) CreateRoomPrices(ctx context.Context, p payloads.CreateRoomPricePayload) error {
	err := uc.checkRoomType(ctx, p.RoomTypeID)
	if err != nil {
		return err
	}
	var rp models.RoomPrice
	err = structs.Merge(&rp, p)
	if err != nil {
		return err
	}

	if p.UntilDate == "" || p.Date == p.UntilDate {
		rp, err = uc.createOneRoomPrice(ctx, rp)
		if err != nil {
			return err
		}
		return nil
	}

	return uc.createManyRoomPrice(ctx, rp, p.UntilDate)

}

func (uc *RoomPriceUsecase) GetRoomTypePriceList(ctx context.Context, typeID int64) ([]models.RoomPrice, error) {
	rts, err := uc.rpRepo.FetchRoomTypePrices(ctx, typeID)
	if err != nil {
		return nil, err
	}
	return rts, nil
}

func (uc *RoomPriceUsecase) UpdateRoomPrice(ctx context.Context, p payloads.UpdateRoomPricePayload, id int64) error {
	var rt models.RoomPrice
	err := structs.Merge(&rt, p)
	if err != nil {
		return err
	}
	err = uc.rpRepo.UpdateRoomPrice(ctx, rt, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *RoomPriceUsecase) DeleteRoomPrice(ctx context.Context, id int64) error {
	return uc.rpRepo.DeleteRoomPrice(ctx, id)
}

func (uc *RoomPriceUsecase) GetRoomPrice(ctx context.Context, id int64) (models.RoomPrice, error) {
	return uc.rpRepo.GetRoomPrice(ctx, id)
}
