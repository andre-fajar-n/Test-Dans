package dans

import (
	"context"
	"dans/entity"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type (
	Job struct {
		Config Config
	}

	Config struct {
		BaseURL string
	}
)

func NewJob(cfg *viper.Viper) entity.JobApi {
	jobConfig := cfg.Sub("third_party").Sub("dans")

	config := Config{
		BaseURL: jobConfig.GetString("base_url"),
	}
	return &Job{
		config,
	}
}

func (j *Job) GetDetail(ctx context.Context, id string) (*entity.JobDetail, error) {
	url := j.Config.BaseURL + fmt.Sprintf("/recruitment/positions/%s", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to call dans API")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := new(entity.JobDetail)
	err = json.Unmarshal(resBody, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (j *Job) GetList(ctx context.Context, form *entity.JobListRequest) (*entity.JobList, error) {
	url := j.Config.BaseURL + "/recruitment/positions.json"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = form.GenerateQueryParam(req)

	fmt.Println("API", req.URL.String())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to call dans API")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := new(entity.JobList)
	err = json.Unmarshal(resBody, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
