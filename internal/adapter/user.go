package adapter

import (
	"card-service/gen/proto"
	"card-service/internal/model"
	"time"
)

func UserToProto(user *model.User) *proto.User {
	return &proto.User{
		Id:           user.ID,
		Name:         user.Name,
		PhoneNumber:  user.PhoneNumber,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    user.UpdatedAt.Format(time.RFC3339),
	}
}

func ProtoToUser(user *proto.User) *model.User {
	createdAt, _ := time.Parse(time.RFC3339, user.GetCreatedAt())
	updatedAt, _ := time.Parse(time.RFC3339, user.GetUpdatedAt())

	return &model.User{
		ID:           user.GetId(),
		Name:         user.GetName(),
		PhoneNumber:  user.GetPhoneNumber(),
		PasswordHash: user.GetPasswordHash(),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
