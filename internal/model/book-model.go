package model

import (
	"context"
	"errors"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-resty/resty"
)

type bookService struct {
	client  *resty.Client
	key     string
	apiRoot string
}

type Book interface {
	Get(ctx context.Context, id string) (vo BookType, err error)
}

type BookConfig struct {
	Client *http.Client
}

// Create will create a new bookService for you, with the given api-key
func (c BookConfig) Create(key string) Book {

	bs := bookService{key: key, apiRoot: "http://www.goodreads.com"}
	if c.Client == nil {
		c.Client = http.DefaultClient
	}
	bs.client = resty.NewWithClient(c.Client)
	bs.client.SetHTTPMode()
	bs.client.SetDebug(false)
	bs.client.SetHostURL("")
	bs.client.SetHeader("Accept", "application/json")

	hystrix.ConfigureCommand("book", hystrix.CommandConfig{
		Timeout:               15000,
		MaxConcurrentRequests: 500,
		ErrorPercentThreshold: 25,
	})

	return bs
}

func (b bookService) Get(ctx context.Context, id string) (BookType, error) {
	response := &Response{}
	err := hystrix.Do("book", func() error {
		res, e := b.client.R().
			SetContext(ctx).
			SetPathParams(map[string]string{
				"id":      id,
				"key":     b.key,
				"apiroot": b.apiRoot,
			}).
			Get("{apiroot}/book/show/{id}.xml?key={key}")
		if res != nil {
			if res.StatusCode() == 404 {
				return errors.New("Not found")
			}
		}
		xmlUnmarshal(res.Body(), response)
		return e
	}, nil)

	return response.Book, err
}
