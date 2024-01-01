package services

import (
	"context"
	"log"
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

func (eventSer *EventServiceImpl) Create(event *models.EventRequest) error {
	id := uuid.New()
	timeNow := time.Now().UTC()
	const layout = "2006-01-02T15:04"

	parsedDate, err := time.Parse(layout, event.EventDetails.Date)
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		return err
	}

	parsedRegistration, err := time.Parse(layout, event.EventDetails.Registration)
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		return err
	}

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
		Date:             parsedDate,
		Registration:     parsedRegistration,
		CreatedAt:        timeNow,
	}

	if err := eventSer.EventRepository.Create(eventModel); err != nil {
		return err
	}

	return nil
}

func (eventSer *EventServiceImpl) Update(id uuid.UUID, event *models.EventRequest) error {
	evt, err := eventSer.EventRepository.FindByID(id)
	if err != nil {
		return err
	}

	eventModel := &models.Event{}
	timeNow := time.Now()

	const layout = "2006-01-02T15:04"

	parsedDate, err := time.Parse(layout, event.EventDetails.Date)
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		return err
	}

	parsedRegistration, err := time.Parse(layout, event.EventDetails.Registration)
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		return err
	}

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
			Date:             parsedDate,
			Registration:     parsedRegistration,
			UpdatedAt:        &timeNow,
		}

	} else {
		eventModel = &models.Event{
			ID:               id,
			Title:            event.EventDetails.Title,
			ShortDescription: event.EventDetails.ShortDescription,
			Description:      event.EventDetails.Description,
			ImageUrl:         *event.EventDetails.ImageUrl,
			ImagePublicId:    *event.EventDetails.ImagePublicId,
			Date:             parsedDate,
			Registration:     parsedRegistration,
			UpdatedAt:        &timeNow,
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

func (eventSer *EventServiceImpl) FindByID(id uuid.UUID) (*models.EventResponse, error) {
	result, err := eventSer.EventRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := &models.EventResponse{
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

func (eventSer *EventServiceImpl) FindAll() ([]*models.EventResponse, error) {
	result, err := eventSer.EventRepository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := []*models.EventResponse{}
	for _, resp_item := range result {
		resp = append(resp, &models.EventResponse{
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
