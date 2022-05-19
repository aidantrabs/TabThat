package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"example/bookmark-api/models"
	
	"context"
	"errors"
)

type BookmarkServiceImpl struct {
	bmcollection *mongo.Collection
	ctx context.Context
}

func _BookmarkService(bmcollection *mongo.Collection, ctx context.Context) BookmarkService {
	return &BMServiceImpl{
		bmcollection: bmcollection,
		ctx: ctx,
	}
}

func (bm *BookmarkServiceImpl) CreateBM(bm *models.Bookmark) error {
	_, err := bmcollection.InsertOne(ctx, bm)
	return err
}

func (bm *BookmarkServiceImpl) GetBM(id *string) (*models.Bookmark, error) {
	var bm models.Bookmark
	err := bmcollection.FindOne(ctx, bm).Decode(&bm)
	return &bm, err
}

func (bm *BookmarkServiceImpl) GetAllBM() ([]*models.Bookmark, error) {
	var bms []*models.Bookmark
	cur, err := bmcollection.Find(ctx, bm)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var bm models.Bookmark
		err := cur.Decode(&bm)
		if err != nil {
			return nil, err
		}
		bms = append(bms, &bm)
	}
	return bms, nil
}

func (bm *BookmarkServiceImpl) UpdateBM(bm *models.Bookmark) error {
	_, err := bmcollection.UpdateOne(ctx, bm, bm)
	return err
}

func (bm *BookmarkServiceImpl) DeleteBM(id *string) error {
	_, err := bmcollection.DeleteOne(ctx, bm)
	return err
}