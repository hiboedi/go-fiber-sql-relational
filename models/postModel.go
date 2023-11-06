package models

type Post struct {
	ID     int          `json:"id" grom:"primaryKey"`
	Title  string       `json:"title" form:"title" grom:"not null"`
	Body   string       `json:"body" form:"body" grom:"not null"`
	UserID int          `json:"user_id" form:"user_id"`
	User   UserResponse `json:"user"`
	Tags   []Tag        `json:"tags" gorm:"many2many:post_tags"`
	TagsID []int        `json:"tags_id" form:"tags_id" gorm:"-"`
}

type PostResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
	UserID int    `json:"user_id" form:"user_id"`
}

type PostResponseHiddenUserID struct {
	ID     int    `json:"id"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
	UserID int    `json:"-" form:"user_id"`
}

type PostResponseWithTag struct {
	ID     int          `json:"id"`
	Title  string       `json:"title" form:"title"`
	Body   string       `json:"body" form:"body"`
	UserID int          `json:"-" form:"user_id"`
	User   UserResponse `json:"user"`
	Tags   []Tag        `json:"tags" gorm:"many2many:post_tags;ForeignKey:ID;joinForeignKey:PostID;Referencs:ID;joinReferences:TagID"`
}

func (PostResponse) TableName() string {
	return "posts"
}

func (PostResponseHiddenUserID) TableName() string {
	return "posts"
}

func (PostResponseWithTag) TableName() string {
	return "posts"
}
