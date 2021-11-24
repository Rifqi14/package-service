package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"gitlab.com/s2.1-backend/shm-package-svc/interfacepkg"
	"strings"
)

type IApi interface {
	Create(index string, id string, body map[string]interface{}) (err error)

	Get(index, ID string) (res map[string]interface{}, err error)

	Update(index, ID string, body map[string]interface{}) (err error)

	Search(index string, query map[string]interface{}) (res map[string]interface{}, err error)
}

type Api struct{
	EsClient *elasticsearch.Client
}

func NewApi(esClient *elasticsearch.Client) IApi {
	return Api{EsClient: esClient}
}

func (api Api) Create(index string, id string, body map[string]interface{}) (err error) {
	bodyStr := interfacepkg.Marshall(body)

	resp, err := api.EsClient.Index(index, strings.NewReader(bodyStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 || resp.IsError() {
		return err
	}

	return nil
}

func (api Api) Get(index, ID string) (res map[string]interface{}, err error) {
	resp, err := api.EsClient.Get(index, ID)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return res, err
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return res, err
		}
	}

	return res, nil
}

func (api Api) Update(index, ID string, body map[string]interface{}) (err error) {
	bodyStr := interfacepkg.Marshall(body)

	resp, err := api.EsClient.Update(index, ID, strings.NewReader(bodyStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 || resp.IsError() {
		return err
	}

	return nil
}

func (api Api) Search(index string, query map[string]interface{}) (res map[string]interface{}, err error) {
	var buffer bytes.Buffer
	err = json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return res, err
	}

	resp, err := api.EsClient.Search(
		api.EsClient.Search.WithContext(context.Background()),
		api.EsClient.Search.WithIndex(index),
		api.EsClient.Search.WithBody(&buffer),
		api.EsClient.Search.WithPretty(),
	)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return res, err
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
