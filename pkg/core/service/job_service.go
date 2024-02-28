package service

import "main/pkg/core/port"

type JobService struct {
	repo port.JobRepository
}

func NewJobService(repo port.JobRepository) *JobService {
	return &JobService{
		repo: repo,
	}
}
