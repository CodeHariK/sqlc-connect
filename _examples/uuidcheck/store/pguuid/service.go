// Code generated by sqlc-connect (https://github.com/walterwanderley/sqlc-connect). DO NOT EDIT.

package pguuid

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "uuidcheck/api/pguuid/v1"
	"uuidcheck/api/pguuid/v1/v1connect"
	"uuidcheck/internal/validation"
)

type Service struct {
	v1connect.UnimplementedPguuidServiceHandler
	querier *Queries
}

func (s *Service) CreateProduct(ctx context.Context, req *connect.Request[pb.CreateProductRequest]) (*connect.Response[pb.CreateProductResponse], error) {
	var arg CreateProductParams
	arg.ID = req.Msg.GetId()
	if v := req.Msg.GetCategory(); v != nil {
		arg.Category = pgtype.Int4{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateProduct(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateProduct")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateProductResponse{Value: result}), nil
}

func (s *Service) CreateProductReturnAll(ctx context.Context, req *connect.Request[pb.CreateProductReturnAllRequest]) (*connect.Response[pb.CreateProductReturnAllResponse], error) {
	var arg CreateProductReturnAllParams
	arg.ID = req.Msg.GetId()
	if v := req.Msg.GetCategory(); v != nil {
		arg.Category = pgtype.Int4{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateProductReturnAll(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateProductReturnAll")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateProductReturnAllResponse{Product: toProduct(result)}), nil
}

func (s *Service) CreateProductReturnPartial(ctx context.Context, req *connect.Request[pb.CreateProductReturnPartialRequest]) (*connect.Response[pb.CreateProductReturnPartialResponse], error) {
	var arg CreateProductReturnPartialParams
	arg.ID = req.Msg.GetId()
	if v := req.Msg.GetCategory(); v != nil {
		arg.Category = pgtype.Int4{Valid: true, Int32: v.Value}
	}

	result, err := s.querier.CreateProductReturnPartial(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateProductReturnPartial")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateProductReturnPartialResponse{CreateProductReturnPartialRow: toCreateProductReturnPartialRow(result)}), nil
}

func (s *Service) CreateUser(ctx context.Context, req *connect.Request[pb.CreateUserRequest]) (*connect.Response[pb.CreateUserResponse], error) {
	var arg CreateUserParams
	if v := req.Msg.GetId(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.ID); err != nil {
			err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}
	if v := req.Msg.GetLocation(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.Location); err != nil {
			err = fmt.Errorf("invalid Location: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}

	result, err := s.querier.CreateUser(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateUser")
		return nil, err
	}
	if uuidStr, err := result.MarshalJSON(); err != nil {
		return nil, fmt.Errorf("failed to convert UUID to string: %w", err)
	} else {
		return connect.NewResponse(&pb.CreateUserResponse{Value: wrapperspb.String(string(uuidStr))}), nil
	}
}

func (s *Service) CreateUserReturnAll(ctx context.Context, req *connect.Request[pb.CreateUserReturnAllRequest]) (*connect.Response[pb.CreateUserReturnAllResponse], error) {
	var arg CreateUserReturnAllParams
	if v := req.Msg.GetId(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.ID); err != nil {
			err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}
	if v := req.Msg.GetLocation(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.Location); err != nil {
			err = fmt.Errorf("invalid Location: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}

	result, err := s.querier.CreateUserReturnAll(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateUserReturnAll")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateUserReturnAllResponse{User: toUser(result)}), nil
}

func (s *Service) CreateUserReturnPartial(ctx context.Context, req *connect.Request[pb.CreateUserReturnPartialRequest]) (*connect.Response[pb.CreateUserReturnPartialResponse], error) {
	var arg CreateUserReturnPartialParams
	if v := req.Msg.GetId(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.ID); err != nil {
			err = fmt.Errorf("invalid ID: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}
	if v := req.Msg.GetLocation(); v != nil {
		if err := json.Unmarshal([]byte(v.GetValue()), &arg.Location); err != nil {
			err = fmt.Errorf("invalid Location: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
	}

	result, err := s.querier.CreateUserReturnPartial(ctx, arg)
	if err != nil {
		slog.Error("sql call failed", "error", err, "method", "CreateUserReturnPartial")
		return nil, err
	}
	return connect.NewResponse(&pb.CreateUserReturnPartialResponse{CreateUserReturnPartialRow: toCreateUserReturnPartialRow(result)}), nil
}
