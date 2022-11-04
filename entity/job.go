package entity

import "context"

// Response
type (
	JobDetail struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		URL         string `json:"url"`
		CreatedAt   string `json:"created_at"`
		Company     string `json:"company"`
		CompanyURL  string `json:"company_url"`
		Location    string `json:"location"`
		Title       string `json:"title"`
		Description string `json:"description"`
		HowToApply  string `json:"how_to_apply"`
		CompanyLogo string `json:"company_logo"`
	}

	JobList []*JobDetail
)

// Request
type (
	JobListRequest struct {
		Description string `query:"description"`
		Location    string `query:"location"`
		FullTime    *bool  `query:"full_time"`
		Page        int    `query:"page"`
	}
)

// Interface
type (
	JobApi interface {
		GetDetail(ctx context.Context, id string) (*JobDetail, error)
		GetList(ctx context.Context, form *JobListRequest) (*JobList, error)
	}

	JobUsecase interface {
		GetDetail(ctx context.Context, id string) (*JobDetail, error)
		GetList(ctx context.Context, form *JobListRequest) (*JobList, error)
	}
)
