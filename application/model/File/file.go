package FileModel

type Resources struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Url      string `json:"url"`
	UserName string `json:"user_name"`
}
