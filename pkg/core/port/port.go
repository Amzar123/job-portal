package port

import (
	"main/pkg/core/domain"

	"github.com/gofiber/fiber/v2"
)

type JobService interface {
	GetJobs(ctx *fiber.Ctx) error
}

type JobRepository interface {
	Get(userID string, message domain.Job) error
}
