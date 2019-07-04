package susi

import (
	"context"
	"net"
	"storage/internal/db"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"google.golang.org/grpc"
)

const pollDelay = 5 * time.Second

type (
	server struct {
		log   Log
		store Storage

		setProduct chan db.Product
		delProduct chan int
	}
	// Storage products to send susi.
	Storage interface {
		GetProducts(ctx Ctx, since time.Time, ch chan<- db.Product) error
	}

	// Configuration contains config for api service.
	Configuration struct {
		Port int
		Host string
	}

	// Ctx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)

func newServer(log Log, s Storage, chSet chan db.Product, chDel chan int) *server {
	return &server{log: log, store: s, setProduct: chSet, delProduct: chDel}
}

func (s *server) GetChanges(since *Since, stream Storage_GetChangesServer) error {
	ctx := stream.Context()
	t := time.Unix(since.UNIX, 0)

	err := s.store.GetProducts(ctx, t, s.setProduct)
	if err != nil {
		return s.log.Err("failed to get tracks", "since", t, "err", err)
	}

	s.log.Info("streaming tracks", "since", t)
	for ; ; time.Sleep(pollDelay) {
		events := dispatcher(s.setProduct, s.delProduct)

		for event := range events {
			if err := stream.Send(event); err != nil {
				return s.log.Err("failed to send track", "err", err)
			}
		}

	}
}

func dispatcher(setProduct <-chan db.Product, delProduct <-chan int) <-chan *ProductEvent {
	events := make(chan *ProductEvent)

	go func() {
		for product := range setProduct {
			convertProduct := convertSet(product)
			events <- &ProductEvent{Event: &ProductEvent_Set{Set: convertProduct}}
		}
	}()

	go func() {
		for productID := range delProduct {
			convertDelProduct := convertDel(productID)
			events <- &ProductEvent{Event: &ProductEvent_Del{Del: convertDelProduct}}
		}
	}()

	return events
}

func convertSet(product db.Product) *SetProduct {
	return &SetProduct{
		Id:          int32(product.ID),
		Brand:       product.Brand,
		Name:        product.Name,
		Description: product.Description,
		Apply:       product.Apply,
		Price:       int32(product.Price),
		Avatar:      product.Avatar.String,
	}
}

func convertDel(productID int) *DelProduct {
	return &DelProduct{Id: int32(productID)}
}

// StartServer - listen grpc server.
func StartServer(log Log, s Storage, chSet chan db.Product, chDel chan int, cfg Configuration) error {
	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	log.Info("gRPC listening", "addr", ln.Addr())
	srv := grpc.NewServer()
	RegisterStorageServer(srv, newServer(log, s, chSet, chDel))

	return errors.Wrap(srv.Serve(ln), "failed to serve")
}
