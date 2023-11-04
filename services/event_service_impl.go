package services

import (
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/repositories"
)

type EventServiceImpl struct {
	EventRepository repositories.EventRepository
	Cloudinary      *cloudinary.Cloudinary
}

func NewEventService(eventRepo repositories.EventRepository, cloudinary *cloudinary.Cloudinary) EventService {
	return &EventServiceImpl{
		EventRepository: eventRepo,
		Cloudinary:      cloudinary,
	}
}

func (eventSer *EventServiceImpl) Create(event *models.EventReq) error {
	id := uuid.New()
	time := time.Now()

	ctx := context.Background()

	result, err := eventSer.Cloudinary.Upload.Upload(ctx, event.ImageFile, uploader.UploadParams{
		PublicID: event.ImageHeader.Filename,
	})

	if err != nil {
		return err
	}

	eventModel := models.Event{
		ID:               id,
		Title:            event.Title,
		ShortDescription: event.ShortDescription,
		Description:      event.Description,
		ImageUrl:         result.SecureURL,
		Date:             event.Date,
		Registration:     event.Registration,
		CreatedAt:        time,
	}

	if err := eventSer.EventRepository.Create(eventModel); err != nil {
		return err
	}

	return nil
}

func (eventSer *EventServiceImpl) Update(id uuid.UUID, event *models.EventReq) error {
	if _, err := eventSer.EventRepository.FindByID(id); err != nil {
		return err
	}

	time := time.Now()
	eventModel := models.Event{
		ID:               id,
		Title:            event.Title,
		ShortDescription: event.ShortDescription,
		Description:      event.Description,
		ImageUrl:         event.ImageUrl,
		Date:             event.Date,
		Registration:     event.Registration,
		UpdatedAt:        &time,
	}

	if err := eventSer.EventRepository.Update(eventModel); err != nil {
		return err
	}
	return nil
}

func (eventSer *EventServiceImpl) Delete(id uuid.UUID) error {
	if _, err := eventSer.EventRepository.FindByID(id); err != nil {
		return err
	}

	if err := eventSer.EventRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (eventSer *EventServiceImpl) FindByID(id uuid.UUID) (*models.EventRes, error) {
	result, err := eventSer.EventRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := &models.EventRes{
		ID:               result.ID,
		Title:            result.Title,
		ShortDescription: result.ShortDescription,
		Description:      result.Description,
		ImageUrl:         result.ImageUrl,
		Date:             result.Date,
		Registration:     result.Registration,
	}

	return resp, nil
}

func (eventSer *EventServiceImpl) FindAll() ([]*models.EventRes, error) {
	result, err := eventSer.EventRepository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := []*models.EventRes{}
	for _, resp_item := range result {
		resp = append(resp, &models.EventRes{
			ID:               resp_item.ID,
			Title:            resp_item.Title,
			ShortDescription: resp_item.ShortDescription,
			Description:      resp_item.Description,
			ImageUrl:         resp_item.ImageUrl,
			Date:             resp_item.Date,
			Registration:     resp_item.Registration,
		})
	}

	return resp, nil

}
