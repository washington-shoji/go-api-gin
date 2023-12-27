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
	time := time.Now().UTC()

	ctx := context.Background()

	result, err := eventSer.Cloudinary.Upload.Upload(ctx, event.ImageFile, uploader.UploadParams{
		PublicID: event.ImageHeader.Filename,
	})

	if err != nil {
		return err
	}

	eventModel := models.Event{
		ID:               id,
		Title:            event.EventDetails.Title,
		ShortDescription: event.EventDetails.ShortDescription,
		Description:      event.EventDetails.Description,
		ImageUrl:         result.SecureURL,
		ImagePublicId:    result.PublicID,
		Date:             event.EventDetails.Date,
		Registration:     event.EventDetails.Registration,
		CreatedAt:        time,
	}

	if err := eventSer.EventRepository.Create(eventModel); err != nil {
		return err
	}

	return nil
}

func (eventSer *EventServiceImpl) Update(id uuid.UUID, event *models.EventReq) error {
	evt, err := eventSer.EventRepository.FindByID(id)
	if err != nil {
		return err
	}

	eventModel := &models.Event{}
	time := time.Now()

	ctx := context.Background()

	if event.ImageFile != nil {
		if _, err := eventSer.Cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{
			PublicID:     evt.ImagePublicId,
			ResourceType: "image",
		}); err != nil {
			return err
		}

		result, err := eventSer.Cloudinary.Upload.Upload(ctx, event.ImageFile, uploader.UploadParams{
			PublicID: event.ImageHeader.Filename,
		})

		if err != nil {
			return err
		}

		eventModel = &models.Event{
			ID:               id,
			Title:            event.EventDetails.Title,
			ShortDescription: event.EventDetails.ShortDescription,
			Description:      event.EventDetails.Description,
			ImageUrl:         result.SecureURL,
			ImagePublicId:    result.PublicID,
			Date:             event.EventDetails.Date,
			Registration:     event.EventDetails.Registration,
			UpdatedAt:        &time,
		}

	} else {
		eventModel = &models.Event{
			ID:               id,
			Title:            event.EventDetails.Title,
			ShortDescription: event.EventDetails.ShortDescription,
			Description:      event.EventDetails.Description,
			ImageUrl:         *event.EventDetails.ImageUrl,
			ImagePublicId:    *event.EventDetails.ImagePublicId,
			Date:             event.EventDetails.Date,
			Registration:     event.EventDetails.Registration,
			UpdatedAt:        &time,
		}
	}

	if err := eventSer.EventRepository.Update(*eventModel); err != nil {
		return err
	}
	return nil
}

func (eventSer *EventServiceImpl) Delete(id uuid.UUID) error {
	evt, err := eventSer.EventRepository.FindByID(id)
	if err != nil {
		return err
	}

	if _, err := eventSer.Cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     evt.ImagePublicId,
		ResourceType: "image",
	}); err != nil {
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
		ImagePublicId:    &result.ImagePublicId,
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
			ImagePublicId:    &resp_item.ImagePublicId,
			Date:             resp_item.Date,
			Registration:     resp_item.Registration,
		})
	}

	return resp, nil

}
