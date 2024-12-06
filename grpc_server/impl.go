package main

import (
	"context"
	"fmt"
	"time"

	"github.com/imran31415/example-project-proto-db/auth"
	"github.com/imran31415/example-project-proto-db/generated_models"
)

// Example gRPC Methods Implementation
func (s *Server) CreateRole(ctx context.Context, req *auth.Role) (*auth.Role, error) {
	role := &generated_models.Role{
		RoleName:  req.GetRoleName(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := role.Insert(ctx, s.Db)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %v", err)
	}

	return &auth.Role{
		RoleId:   int32(role.RoleID),
		RoleName: role.RoleName,
	}, nil
}

// Example gRPC Methods Implementation
func (s *Server) CreateUser(ctx context.Context, req *auth.User) (*auth.User, error) {

	user := &generated_models.User{
		UserID:    int(req.GetUserId()),
		Username:  req.Username,
		Email:     req.GetEmail(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := user.Insert(ctx, s.Db)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %v", err)
	}

	return req, nil
}

func (s *Server) GetRoleById(ctx context.Context, req *auth.GetRoleRequest) (*auth.Role, error) {
	role, err := generated_models.RoleByRoleID(ctx, s.Db, int(req.GetRoleId()))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve role: %v", err)
	}

	return &auth.Role{
		RoleId:   int32(role.RoleID),
		RoleName: role.RoleName,
	}, nil
}

func (s *Server) UpdateRole(ctx context.Context, req *auth.Role) (*auth.Role, error) {
	role, err := generated_models.RoleByRoleID(ctx, s.Db, int(req.GetRoleId()))
	if err != nil {
		return nil, fmt.Errorf("failed to find role: %v", err)
	}

	role.RoleName = req.GetRoleName()
	role.UpdatedAt = time.Now()

	err = role.Update(ctx, s.Db)
	if err != nil {
		return nil, fmt.Errorf("failed to update role: %v", err)
	}

	return &auth.Role{
		RoleId:   int32(role.RoleID),
		RoleName: role.RoleName,
	}, nil
}

func (s *Server) DeleteRole(ctx context.Context, req *auth.Role) (*auth.Role, error) {
	role, err := generated_models.RoleByRoleID(ctx, s.Db, int(req.GetRoleId()))
	if err != nil {
		return nil, fmt.Errorf("failed to find role: %v", err)
	}

	err = role.Delete(ctx, s.Db)
	if err != nil {
		return nil, fmt.Errorf("failed to delete role: %v", err)
	}

	return &auth.Role{}, nil
}
