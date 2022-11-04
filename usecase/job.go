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

func (u *Job) GetList(ctx context.Context, form *entity.JobListRequest) (*entity.JobList, error) {
	data, err := u.jobApi.GetList(ctx, form)
	if err != nil {
		return nil, err
	}

	output := entity.JobList{}
	for _, v := range *data {
		if v != nil {
			output = append(output, v)
		}
	}

	return &output, nil
}
