package services

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/repositories"
)

type DynamicDataServiceImpl struct {
	Repository repositories.DynamicDataRepository
}

func NewDynamicDataService(repo repositories.DynamicDataRepository) DynamicDataService {
	return &DynamicDataServiceImpl{
		Repository: repo,
	}
}

// Create implements DynamicDataService.
func (dynRepo *DynamicDataServiceImpl) Create(data *models.DynamicDataReq) error {
	id := uuid.New()
	dynDataModel := &models.DynamicData{
		ID:   id,
		Data: data.Data,
	}

	err := dynRepo.Repository.Create(dynDataModel)
	if err != nil {
		return err
	}

	return nil
}

// Update implements DynamicDataService.
func (dynRepo *DynamicDataServiceImpl) Update(id uuid.UUID, data *models.DynamicDataReq) error {
	dynDataModel := &models.DynamicData{
		Data: data.Data,
	}

	err := dynRepo.Repository.Update(id, dynDataModel)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements DynamicDataService.
func (dynRepo *DynamicDataServiceImpl) Delete(id uuid.UUID) error {
	err := dynRepo.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements DynamicDataService.
func (dynRepo *DynamicDataServiceImpl) FindAll() ([]*models.DynamicDataRes, error) {
	result, err := dynRepo.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := []*models.DynamicDataRes{}

	for _, result_item := range result {
		resp = append(resp, &models.DynamicDataRes{
			ID:   result_item.ID,
			Data: result_item.Data,
		})
	}

	return resp, nil
}

// FindById implements DynamicDataService.
func (dynRepo *DynamicDataServiceImpl) FindById(id uuid.UUID) (*models.DynamicDataRes, error) {
	result, err := dynRepo.Repository.FindById(id)
	if err != nil {
		return nil, err
	}

	resp := &models.DynamicDataRes{
		ID:   result.ID,
		Data: result.Data,
	}

	return resp, nil
}
