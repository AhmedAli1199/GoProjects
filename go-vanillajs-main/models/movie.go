package models

type Movie struct{
	ID int
	TMBD_ID int
	Title string
	Tagline string
	ReleaseYear int
	Genres []Genre
	Overview *string
	Score *float32
	Popularitey *float32
	Keywords []string
	Language *string
	PosterURL *string
	TrailerURL *string
	Casting []Actor

}