package handlers

import (
	"context"
	"time"

	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/models"
	"github.com/gofiber/fiber/v2"
	
)
	type EventHandler struct{
		repository models.EventRepository
	}

func (h *EventHandler) GetMany(ctx *fiber.Ctx) error{
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	events, err := h.repository.GetMany(context)

	if err != nil{
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status": "fail",
            "message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		    "status": "success",
			"message" : "",
            "data": events,
        })
}

func (h *EventHandler) GetOne(ctx *fiber.Ctx) error{
	context, cancel:= context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	eventId := ctx.Params("id")
	events, err:= h.repository.GetOne(context, eventId)
	if err!= nil{
		// if err == models.ErrNotFound{
		// 	return ctx.Status(fiber.StatusNotFound).JSON(&fiber.Map{
        //         "status":  "fail",
        //         "message": "Event not found",
        //     })
        // }

		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
            "status": "fail",
            "message": err.Error(),
        })
	}
	
    return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		    "status": "success",
            "message": "",
            "data": events,
        })
}
// func (h *EventHandler) CreateOne(ctx fiber.Ctx) error{
// 	context, cancel:= context.WithTimeout(context.Background(), time.Duration(5*time.Second))
// 	defer cancel()
// 	var event models.Event
// 	err:= ctx.BodyParser(&event)
// 	if err!= nil{
// 		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
//             "status":  "fail",
//             "message": "Invalid request body",
//         })
    
// }
func NewEventHandler(router fiber.Router,repository models.EventRepository){
	handler:= &EventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	// router.Post("/", handler.CreateOne)
	router.Get("/:id", handler.GetOne)
} 