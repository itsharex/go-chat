package model

type GroupNotice struct {
	ID           int    `json:"id" grom:"comment:群公告ID"`
	GroupId      int    `json:"group_id" grom:"comment:群组ID"`
	CreatorId    int    `json:"creator_id" grom:"comment:创建者用户ID"`
	Title        string `json:"title" grom:"comment:公告标题"`
	Content      string `json:"content" grom:"comment:公告内容"`
	IsTop        int    `json:"is_top" grom:"comment:是否置顶"`
	IsDelete     int    `json:"is_delete" grom:"comment:是否删除"`
	IsConfirm    int    `json:"is_confirm" grom:"comment:是否需群成员确认公告"`
	ConfirmUsers string `json:"confirm_users" grom:"comment:已确认成员"`
	CreatedAt    string `json:"created_at" grom:"comment:创建时间"`
	UpdatedAt    string `json:"updated_at" grom:"comment:更新时间"`
	DeletedAt    string `json:"deleted_at" grom:"comment:删除时间"`
}