package model

import (
	"context"
	"errors"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-resty/resty"
)

type userService struct {
	client  *resty.Client
	key     string
	apiRoot string
}

type User interface {
	Get(ctx context.Context, id string) (vo UserType, err error)
}

type UserConfig struct {
	Client *http.Client
}

// Create will create a new bookService for you, with the given api-key
func (c UserConfig) Create(key string) User {

	us := userService{key: key, apiRoot: "http://www.goodreads.com"}
	if c.Client == nil {
		c.Client = http.DefaultClient
	}
	us.client = resty.NewWithClient(c.Client)
	us.client.SetHTTPMode()
	us.client.SetDebug(false)
	us.client.SetHostURL("")
	us.client.SetHeader("Accept", "application/json")

	hystrix.ConfigureCommand("user", hystrix.CommandConfig{
		Timeout:               15000,
		MaxConcurrentRequests: 500,
		ErrorPercentThreshold: 25,
	})

	return us
}

func (b userService) Get(ctx context.Context, id string) (UserType, error) {

	response := &Response{}
	err := hystrix.Do("user", func() error {
		res, e := b.client.R().
			SetContext(ctx).
			SetPathParams(map[string]string{
				"id":      id,
				"key":     b.key,
				"apiroot": b.apiRoot,
			}).
			Get("{apiroot}/user/show/{id}.xml?key={key}")
		if res != nil {
			if res.StatusCode() == 404 {
				return errors.New("Not found")
			}
		}
		xmlUnmarshal(res.Body(), response)
		return e
	}, nil)

	// for i := range response.User.Statuses {
	// 	status := &response.User.Statuses[i]
	// 	bookid := status.Book.ID
	// 	book := GetBook(bookid, key)
	// 	status.Book = book
	// }
	// limit := 10
	// if len(response.User.Statuses) >= limit {
	// 	response.User.Statuses = response.User.Statuses[:limit]
	// } else {
	// 	remaining := limit - len(response.User.Statuses)
	// 	response.User.LastRead = GetLastRead(id, key, remaining)
	// }

	return response.User, err
}
