package pony

import (
	"context"
	"net/url"
)

type Database interface {
	InsertMessage(
		ctx context.Context,
		topic string,
		contentType string,
		data []byte,
		queryValues url.Values,
	) error
}
