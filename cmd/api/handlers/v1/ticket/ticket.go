package ticket

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
	"strconv"
	"ticket-booking-app/internal/db/models"
	"ticket-booking-app/internal/db/repository"
)

var _ Handler = (*Ticket)(nil)

func NewTicketHandler(repo repository.Ticket) *Ticket {
	return &Ticket{
		repo: repo,
	}
}

type Ticket struct {
	repo repository.Ticket
}

func (t Ticket) GetMany(c *fiber.Ctx) error {
	userId := uint(c.Locals("userId").(float64))

	tickets, err := t.repo.GetMany(context.Background(), userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    tickets,
	})
}

func (t Ticket) GetOne(c *fiber.Ctx) error {
	ticketId, _ := strconv.Atoi(c.Params("ticketId"))
	userId := uint(c.Locals("userId").(float64))

	ticket, err := t.repo.GetOne(context.Background(), userId, uint(ticketId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	var QRCode []byte
	QRCode, err = qrcode.Encode(
		fmt.Sprintf("ticketId:%v,ownerId:%v", ticketId, userId),
		qrcode.Medium,
		256,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data": &fiber.Map{
			"ticket": ticket,
			"qrcode": QRCode,
		},
	})
}

func (t Ticket) CreateOne(c *fiber.Ctx) error {
	ticket := &models.Ticket{}
	userId := uint(c.Locals("userId").(float64))

	if err := c.BodyParser(ticket); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	ticket, err := t.repo.CreateOne(context.Background(), userId, ticket)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Ticket created",
		"data":    ticket,
	})
}

func (t Ticket) UpdateOne(c *fiber.Ctx) error {
	validateBody := &models.ValidateTicket{}

	if err := c.BodyParser(validateBody); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true

	ticket, err := t.repo.UpdateOne(context.Background(), validateBody.OwnerId, validateBody.TicketId, validateData)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Welcome to the show!",
		"data":    ticket,
	})
}
