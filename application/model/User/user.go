package UserModel

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	CreateAt int64  `json:"create_at" form:"create_at"`
	Avatar   string `json:"avatar" form:"avatar"`
	Slogan   string `json:"slogan" form:"slogan"`
}
