package services

import (
	"example/bookmark-api/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"
	"errors"
)

type BookmarkServiceImpl struct {
	bmcollection *mongo.Collection
	ctx context.Context
}

func NewBookmarkService(bmcollection *mongo.Collection, ctx context.Context) BookmarkService {
	return &BookmarkServiceImpl{
		bmcollection: bmcollection,
		ctx: ctx,
	}
}

func (bms *BookmarkServiceImpl) CreateBM(bm *models.Bookmark) error {
	_, err := bms.bmcollection.InsertOne(bms.ctx, bm)
	if err != nil {
		return err
	}
	return nil
}

func (bms *BookmarkServiceImpl) GetBM(id *string) (*models.Bookmark, error) {
	var bm *models.Bookmark
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := bms.bmcollection.FindOne(bms.ctx, query).Decode(&bm)

	if err != nil {
		return nil, err
	}

	return bm, nil
}

func (bms *BookmarkServiceImpl) GetAllBM() ([]*models.Bookmark, error) {
	var bookmarks []*models.Bookmark
	cursor, err := bms.bmcollection.Find(bms.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(bms.ctx) {
		var bookmark models.Bookmark
		if cursor.Decode(&bookmark) != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, &bookmark)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(bms.ctx)

	if (len(bookmarks) == 0) {
		return nil, errors.New("No bookmarks found")
	}

	return bookmarks, nil
}

func (bms *BookmarkServiceImpl) UpdateBM(bm *models.Bookmark) error {

	item, _ := bms.bmcollection.UpdateOne(
		bms.ctx, 
		bson.D{primitive.E{Key: "id", Value: bm.ID}}, 
		bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "id", Value: bm.ID}, primitive.E{Key: "name", Value: bm.Name}, primitive.E{Key: "url", Value: bm.URL}, primitive.E{Key: "description", Value: bm.Description}}}})

	if item.MatchedCount != 1 {
		return errors.New("No bookmark to update")
	}
	return nil
}

func (bms *BookmarkServiceImpl) DeleteBM(id *string) error {

	item, _ := bms.bmcollection.DeleteOne(
		bms.ctx, 
		bson.D{primitive.E{Key: "id", Value: id}})

	if item.DeletedCount != 1 {
		return errors.New("No bookmark to delete")
	}

	return nil
}