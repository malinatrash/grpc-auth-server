package validation

import (
	ssov1 "github.com/malinatrash/grpc-auth-protos/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyValue  = 0
	emptyString = ""
)

func ValidateEmailAndPassword(email string, password string) error {
	if email == emptyString {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if password == emptyString {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}

func ValidateLogin(req *ssov1.LoginRequest) error {
	if err := ValidateEmailAndPassword(req.GetEmail(), req.GetPassword()); err != nil {
		return err
	}
	if req.AppId == emptyValue {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}
	return nil
}

func ValidateRegister(req *ssov1.RegisterRequest) error {
	if err := ValidateEmailAndPassword(req.GetEmail(), req.GetPassword()); err != nil {
		return err
	}
	return nil
}

func ValidateIsAdmin(req *ssov1.IsAdminRequest) error {
	if req.UserId == emptyValue {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	return nil
}
