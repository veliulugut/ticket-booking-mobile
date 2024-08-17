package event

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"ticket-booking-app/internal/db/models"
	"ticket-booking-app/internal/db/repository"
	"time"
)

var _ Handler = (*Event)(nil)

func NewEventHandler(repo repository.Event) *Event {
	return &Event{
		repo: repo,
	}
}

type Event struct {
	repo repository.Event
}

func (e Event) GetMany(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	events, err := e.repo.GetMany(context)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    events,
	})
}

func (e Event) GetByID(c *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(c.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	event, err := e.repo.GetByID(context, strconv.Itoa(int(uint(eventId))))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    event,
	})
}

func (e Event) CreateEvent(c *fiber.Ctx) error {
	event := &models.Event{}

	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	res, err := e.repo.CreateEvent(context.Background(), event)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Event created successfully",
		"data":    res,
	})

}

func (e Event) UpdateEvent(c *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(c.Params("eventId"))
	updateData := make(map[string]interface{})

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	event, err := e.repo.UpdateEvent(context, uint(eventId), updateData)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Event updated",
		"data":    event,
	})
}

func (e Event) DeleteEvent(c *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(c.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := e.repo.DeleteEvent(context, uint(eventId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
