package handler

import (
	"main/pkg/core/service"

	"github.com/gofiber/fiber/v2"
)

type JobHandler struct {
	svc service.JobService
}

func NewJobHandler(svc service.JobService) *JobHandler {
	return &JobHandler{
		svc: svc,
	}
}

// Path: pkg/adapters/handler/job_handler.go
// @id job.Handler.GetJobs
// @summary Get a company
// @tags company
// @produce json
// @router /jobs [get]
// @param field_op query string false "Filter by <field> with <op> is operator"
// @param with query string false "Get related resources"
// @param sort query string false "Sort by <field> in ascending order, add dash (-) in front of <field> to sort in descending order"
// @param skip query int false "Skip the first <skip> records"
// @param limit query int false "Limit result to <limit> records"
// @success 200
func (h *JobHandler) GetJobs(ctx *fiber.Ctx) error {
	return nil
}
