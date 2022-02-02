package TypeModel

type Types struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Type   int    `json:"type"`
	UserId int    `json:"user_id"`
}
