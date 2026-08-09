package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/config"
	GamePromo "github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/api/grpc/proto"
	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/entities"
	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/repository"
	logger "github.com/AlekseiYuzhanin/GamePromoLogger"
	"google.golang.org/grpc"
)

type Server struct {
	conn           *grpc.Server
	ln             net.Listener
	log            logger.Logger
	itemRepository repository.ItemsRepository
	GamePromo.UnimplementedItemsServiceServer
}

func New(cfg *config.Config, log logger.Logger, itemsRepository repository.ItemsRepository) (*Server, error) {
	srv := grpc.NewServer()
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcServer.Port))
	if err != nil {
		return nil, err
	}
	impl := &Server{
		conn:           srv,
		ln:             ln,
		log:            log,
		itemRepository: itemsRepository,
	}
	GamePromo.RegisterItemsServiceServer(srv, impl)
	return impl, nil
}

func (s *Server) CreateItem(ctx context.Context, item *GamePromo.Item) (*GamePromo.Item, error) {
	if item.Type == nil {
		return nil, errors.New("не указан тип слота")
	}
	if item.Rarity == nil {
		return nil, errors.New("не указана редкость слота")
	}
	slot := &entities.Slot{
		Title: item.Title,
		Type: entities.SlotType{
			Title: item.Type.GetTitle(),
		},
		Rarity: entities.SlotRarity{
			Title:      item.Rarity.GetTitle(),
			Weight:     float64(item.Rarity.GetWeight()),
			Multiplier: float64(item.Rarity.GetMultiplier()),
			Color:      item.Rarity.GetColor(),
		},
	}
	_, err := s.itemRepository.CreateSlot(ctx, slot)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Server) CreateItems(ctx context.Context, items *GamePromo.Items) (*GamePromo.Items, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateItem(ctx context.Context, request *GamePromo.UpdateItemRequest) (*GamePromo.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateItems(ctx context.Context, request *GamePromo.UpdateItemsRequest) (*GamePromo.Items, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Start() error {
	return s.conn.Serve(s.ln)
}

func (s *Server) Close() {
	s.conn.GracefulStop()
	if err := s.ln.Close(); err != nil {
		s.log.Error("Error close connection", logger.Field{Key: "Err", Value: err})
	}
	s.log.Info("GRPC server closed...")
}
