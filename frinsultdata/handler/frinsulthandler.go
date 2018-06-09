package handler

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/pijalu/go.hands.two/frinsultdata/model"
	"github.com/pijalu/go.hands.two/frinsultdata/repository"
	"github.com/pijalu/go.hands.two/frinsultproto"
	"github.com/pkg/errors"
)

// FrinsultHandler is the RPC implementation
type FrinsultHandler struct{}

// GetFrinsultByID finds an insult by id
func (f *FrinsultHandler) GetFrinsultByID(c context.Context, req *frinsultproto.ByIDRequest, rep *frinsultproto.Frinsult) error {
	ID := uint(req.GetID())
	if ID == 0 {
		return errors.New("invalid ID")
	}

	dbInsult, err := repository.GetFrinsultByID(ID)
	if err != nil {
		return err
	}

	*rep = frinsultproto.Frinsult{
		ID:    int64(dbInsult.ID),
		Text:  dbInsult.Text,
		Score: int64(dbInsult.Score),
	}
	return nil
}

// DeleteFrinsultByID deletes an insults by id
func (f *FrinsultHandler) DeleteFrinsultByID(c context.Context, req *frinsultproto.ByIDRequest, _ *frinsultproto.Void) error {
	ID := uint(req.GetID())
	if ID == 0 {
		return errors.New("invalid ID")
	}

	return repository.DeleteFrinsultByID(ID)
}

// UpdateFrinsult updates an insult
func (f *FrinsultHandler) UpdateFrinsult(c context.Context, req *frinsultproto.Frinsult, _ *frinsultproto.Void) error {
	ID := uint(req.GetID())
	if ID == 0 {
		return errors.New("invalid ID")
	}

	return repository.UpdateFrinsult(&model.Frinsult{
		Model: gorm.Model{ID: ID},
		Text:  req.GetText(),
		Score: int(req.GetScore()),
	})
}

// VoteFrinsultByID update score of a given insult
func (f *FrinsultHandler) VoteFrinsultByID(c context.Context, v *frinsultproto.VoteRequest, _ *frinsultproto.Void) error {
	ID := uint(v.GetID())
	if ID == 0 {
		return errors.New("invalid ID")
	}

	return repository.VoteForFrinsult(uint(v.GetID()), int(v.GetVote()))
}

// GetFrinsults returns a list of items
func (f *FrinsultHandler) GetFrinsults(c context.Context, _ *frinsultproto.Void, s frinsultproto.FrinsultService_GetFrinsultsStream) error {
	insults, err := repository.GetFrinsults()
	if err != nil {
		return err
	}

	for _, fri := range insults {
		if err := s.SendMsg(&fri); err != nil {
			return errors.Wrap(err, "Error streaming records")
		}
	}

	return nil
}
