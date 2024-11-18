package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-books/services/books"
	"github.com/kaellybot/kaelly-books/utils/insights"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	booksService books.Service
	broker       amqp.MessageBroker
	probes       insights.Probes
	prom         insights.PrometheusMetrics
}
