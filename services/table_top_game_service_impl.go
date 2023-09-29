package services

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/repositories"
)

type TableTopGameServiceImpl struct {
	Repository repositories.TableTopGameRepository
}

func NewTableTopGameService(tgmRep repositories.TableTopGameRepository) TableTopGameService {
	return &TableTopGameServiceImpl{
		Repository: tgmRep,
	}
}

// Create implements TableTopGameService.
func (tgmRep *TableTopGameServiceImpl) Create(game *models.TableTopGameReq) error {

	id := uuid.New()
	tblGameModel := &models.TableTopGame{
		ID:         id,
		Name:       game.Name,
		GameDetail: game.GameDetail,
	}

	err := tgmRep.Repository.Create(tblGameModel)
	if err != nil {
		return err
	}

	return nil
}

// Update implements TableTopGameService.
func (tgmRep *TableTopGameServiceImpl) Update(id uuid.UUID, game *models.TableTopGameReq) error {
	tblGameModel := &models.TableTopGame{
		Name:       game.Name,
		GameDetail: game.GameDetail,
	}

	err := tgmRep.Repository.Update(id, tblGameModel)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements TableTopGameService.
func (tgmRep *TableTopGameServiceImpl) Delete(id uuid.UUID) error {
	err := tgmRep.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements TableTopGameService.
func (tgmRep *TableTopGameServiceImpl) FindAll() ([]*models.TableTopGameResp, error) {
	result, err := tgmRep.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := []*models.TableTopGameResp{}

	for _, result_item := range result {
		resp = append(resp, &models.TableTopGameResp{
			ID:         result_item.ID,
			Name:       result_item.Name,
			GameDetail: result_item.GameDetail,
		})
	}

	return resp, nil
}

// FindByID implements TableTopGameService.
func (tgmRep *TableTopGameServiceImpl) FindByID(id uuid.UUID) (*models.TableTopGameResp, error) {
	result, err := tgmRep.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := &models.TableTopGameResp{
		ID:         result.ID,
		Name:       result.Name,
		GameDetail: result.GameDetail,
	}

	return resp, nil
}
