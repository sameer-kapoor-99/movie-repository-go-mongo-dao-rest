package models

type Movie struct {
	ID string `json:"id" bson:"_id,omitempty"`
	ImdbId string `json:"imdbid" bson:"_imdbid,omitempty"`
	Title string `json:"title" bson:"_title,omitempty"`
	Rating float32 `json:"rating,omitempty"`
	Cast []string `json:"cast,omitempty"`
}
