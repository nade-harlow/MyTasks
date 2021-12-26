package models

type Task struct {
	Title        string `json:"title" bson:"title"`
	Done         bool   `json:"done" bson:"done"`
	CreatedAt    string `json:"created_at" bson:"created_at"`
	LastModified string `json:"last_modified" bson:"last_modified"`
}
