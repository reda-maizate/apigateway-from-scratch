package v1

import (
	business "api-gateway/internal/business/users"
	userstubs "api-gateway/stubs/go/apigateway-from-scratch/users/v1"
	"context"
)

type UserServiceServer struct {
	userBusiness business.UsersBusiness
	userstubs.UnimplementedUserServer
}

func NewUserServiceServer(userBusiness *business.UsersBusiness) userstubs.UserServer {
	return &UserServiceServer{userBusiness: *userBusiness}
}

func (s *UserServiceServer) SignUp(ctx context.Context, req *userstubs.SignUpRequest) (*userstubs.UserResponse, error) {
	signUpParams := business.SignupParams{
		Email:    req.Email,
		Password: req.Password,
	}

	userResponse, err := s.userBusiness.SignUp(ctx, signUpParams)
	if err != nil {
		return nil, err
	}

	return &userstubs.UserResponse{Token: userResponse.Token}, nil
}

func (s *UserServiceServer) Login(ctx context.Context, req *userstubs.LoginRequest) (*userstubs.UserResponse, error) {
	logInParams := business.LoginParams{
		Email:    req.Email,
		Password: req.Password,
	}

	userResponse, err := s.userBusiness.Login(ctx, logInParams)
	if err != nil {
		return nil, err
	}

	return &userstubs.UserResponse{Token: userResponse.Token}, nil
}

func (s *UserServiceServer) UserFromToken(ctx context.Context, req *userstubs.MeUserRequest) (*userstubs.MeUserResponse, error) {
	userFromTokenParams := business.UserFromTokenParams{
		Token: req.Token,
	}
	userFromTokenResponse, err := s.userBusiness.UserFromToken(ctx, userFromTokenParams)
	if err != nil {
		return nil, err
	}

	return &userstubs.MeUserResponse{Id: userFromTokenResponse.Uuid}, nil
}
