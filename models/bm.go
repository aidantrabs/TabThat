package models

type Bookmark struct {
	ID 			int 		`json:"id" bson:"_id"` 
	Name 		string 	`json:"name" bson:"name"`
	URL 			string 	`json:"url" bson:"url"`
	Description 	string 	`json:"description" bson:"description"`
}