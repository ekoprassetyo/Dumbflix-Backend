package filmdto

type CreateFilmRequest struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          string `json:"year" gorm:"type: varchar(255)"`
	Description   string `json:"description" gorm:"type: text"`
}

type UpdateFilmRequest struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          string `json:"year" gorm:"type: varchar(255)"`
	Description   string `json:"description" gorm:"type: text"`
}