package service

import (
	"context"
	"errors"
)

func (s Service) RestoreDeleted(ctx context.Context, uuid string) error {
	hotels, err := s.repo.FindByUUID(ctx, uuid)
	if err != nil {
		return err
	}
	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	userUUID := s.getUserUUIDInAuth(ctx)
	if userUUID != hotels.CreatedBy {
		return errors.New("you are not the creator of this hotels")
	}
	hotels.DeletedAt = nil
	if err = s.repo.Update(ctx, *hotels); err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
