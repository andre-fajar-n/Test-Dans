package usecase

import (
	"context"
	"dans/entity"
)

type (
	Job struct {
		jobApi entity.JobApi
	}
)

func NewJob(jobApi entity.JobApi) entity.JobUsecase {
	return &Job{
		jobApi,
	}
}

func (u *Job) GetDetail(ctx context.Context, id string) (*entity.JobDetail, error) {
	return u.jobApi.GetDetail(ctx, id)
}
