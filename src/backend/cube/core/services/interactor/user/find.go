package user

import (
	"context"
	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/pkg/logger"
)

func (i *Interactor) Find(ctx context.Context, filter interactors.UserInfoFilter) ([]user.Info, error) {
	lg := i.lg.WithFields(map[string]interface{}{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]interface{}{"filter": filter}).Info("user find called")
	return i.userInfo.Find(
		repositories.UserInfoFilter{
			UserIds:   filter.Ids,
			Usernames: filter.Usernames,
			Emails:    filter.Emails,
			Limit:     filter.Limit,
			Offset:    filter.Offset,
		},
	)
}
