package models

type Film struct {
	ID            int                `json:"id"					gorm:"primary_key:auto_increment"`
	Title         string             `json:"title"				gorm:"type: varchar(255)"`
	ThumbnailFilm string             `json:"thumbnailfilm"		gorm:"type: varchar(255)"`
	Year          string             `json:"year"				gorm:"type: varchar(255)"`
	CategoryID    int                `json:"category_id" form:"category_id"`
	Category      CategoriesResponse `json:"category"`
	Description   string             `json:"description"		gorm:"type: text"`
}

type FilmResponse struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	ThumbnailFilm string             `json:"thumbnailfilm"`
	Year          string             `json:"year"`
	CategoryID    int                `json:"-"`
	Category      CategoriesResponse `json:"category"`
	Description   string             `json:"description"`
}

type FilmEpisodeResponse struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	ThumbnailFilm string             `json:"thumbnailfilm"`
	Year          string             `json:"year"`
	CategoryID    int                `json:"-"`
	Category      CategoriesResponse `json:"category"`
}

type FilmCategoryResponse struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	ThumbnailFilm string             `json:"thumbnailfilm"`
	Year          string             `json:"year"`
	CategoryID    int                `json:"-"`
	Category      CategoriesResponse `json:"category"`
	Description   string             `json:"description"`
}

func (FilmCategoryResponse) TableName() string {
	return "films"
}

func (FilmEpisodeResponse) TableName() string {
	return "films"
}

func (FilmResponse) TableName() string {
	return "films"
}
