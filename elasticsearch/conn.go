package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v7"
	"net/http"
	"time"
)

type Conn struct {
	Address  []string
	UserName string
	Password string
}

func NewCon(address []string, userName, password string) (res *elasticsearch.Client, err error) {
	esCon := Conn{
		Address:  address,
		UserName: userName,
		Password: password,
	}

	return esCon.Connection()
}

func (con Conn) Connection() (res *elasticsearch.Client, err error) {
	config := elasticsearch.Config{
		Addresses: con.Address,
		Username:  con.UserName,
		Password:  con.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 3 * time.Second,
		},
	}
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		return res, err
	}

	return es, nil
}
