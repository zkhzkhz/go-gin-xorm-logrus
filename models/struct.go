package models

import (
	"time"
)

type Comment struct {
	Id          string    `xorm:"not null pk VARCHAR(255)"`
	PostId      string    `xorm:"not null index VARCHAR(255)"`
	UserId      string    `xorm:"not null comment('评论用户id') index VARCHAR(255)"`
	FloorNumber int       `xorm:"not null INT(11)"`
	ParentId    string    `xorm:"comment('父评论id') index VARCHAR(255)"`
	ReplyUserId string    `xorm:"comment('评论回复用户id') index VARCHAR(255)"`
	Context     string    `xorm:"TEXT"`
	IsAdmin     int       `xorm:"comment('是否管理员回复；1普通，2管理员') INT(11)"`
	Status      int       `xorm:"not null default 1 comment('状态；-1已删除, 1正常') INT(11)"`
	CreateDate  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

type Contribute struct {
	Id         string    `xorm:"not null pk VARCHAR(255)"`
	PostTypeId string    `xorm:"not null index VARCHAR(255)"`
	UserId     string    `xorm:"not null index VARCHAR(255)"`
	Title      string    `xorm:"VARCHAR(255)"`
	Context    string    `xorm:"TEXT"`
	Examine    int       `xorm:"not null default 1 comment('审核状态；1：待审核，2：审核通过，3：审核不通过') INT(11)"`
	Status     int       `xorm:"not null default 1 comment('状态；-1已删除, 1正常') INT(11)"`
	CreateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Gold       int       `xorm:"INT(11)"`
	Objection  string    `xorm:"comment('拒绝理由') VARCHAR(2000)"`
}

type Group struct {
	Id   string `xorm:"not null pk VARCHAR(255)"`
	Name string `xorm:"VARCHAR(255)"`
}

type KsAdvertisement struct {
	Id          string    `xorm:"not null pk VARCHAR(255)"`
	AdsName     string    `xorm:"not null comment('广告名称') VARCHAR(255)"`
	Url         string    `xorm:"comment('内容链接') VARCHAR(255)"`
	Duration    int       `xorm:"not null comment('广告时长') INT(11)"`
	RedirectUrl string    `xorm:"comment('跳转链接') VARCHAR(255)"`
	Status      int       `xorm:"not null comment('状态（启用：1 停用：0）') INT(11)"`
	ActiveDate  time.Time `xorm:"comment('启用时间') DATETIME"`
	CreateDate  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('配置时间') TIMESTAMP"`
	Type        int       `xorm:"not null comment('类型（图片：1 视频：0）') INT(11)"`
	Attr1       string    `xorm:"VARCHAR(255)"`
	Attr2       string    `xorm:"VARCHAR(255)"`
	Attr3       string    `xorm:"VARCHAR(255)"`
}

type KsFeedback struct {
	Id         string    `xorm:"not null pk VARCHAR(255)"`
	UserId     string    `xorm:"not null comment('用户ID') VARCHAR(255)"`
	PicUrl     string    `xorm:"comment('反馈图片地址') VARCHAR(255)"`
	Content    string    `xorm:"comment('反馈内容') VARCHAR(255)"`
	CreateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('提交时间') TIMESTAMP"`
	Status     int       `xorm:"not null comment('状态（处理：1 未处理：0 忽略：-1）') INT(11)"`
	Attr1      string    `xorm:"VARCHAR(255)"`
	Attr2      string    `xorm:"VARCHAR(255)"`
	Attr3      string    `xorm:"VARCHAR(255)"`
}

type KsHelp struct {
	Id          string    `xorm:"not null pk VARCHAR(255)"`
	Title       string    `xorm:"comment('标题') VARCHAR(255)"`
	ReleaseDate time.Time `xorm:"comment('配置时间') DATETIME"`
	CreateDate  time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	Status      int       `xorm:"not null comment('状态（启用：1 停用：-1）') INT(11)"`
	Content     string    `xorm:"not null comment('内容') VARCHAR(255)"`
	Attr1       string    `xorm:"VARCHAR(255)"`
	Attr2       string    `xorm:"VARCHAR(255)"`
	Attr3       string    `xorm:"VARCHAR(255)"`
}

type KsIconTheme struct {
	Id                     string    `xorm:"not null pk VARCHAR(255)"`
	ThemeTitle             string    `xorm:"comment('主题名称') VARCHAR(255)"`
	ActiveTime             time.Time `xorm:"comment('启用时间') TIMESTAMP"`
	CreateTime             time.Time `xorm:"comment('配置时间') TIMESTAMP"`
	Status                 int       `xorm:"not null comment('状态（启用：1 停用：0）') INT(11)"`
	AppIconSelected        string    `xorm:"comment('app图标 ') VARCHAR(255)"`
	SowingIconUnselected   string    `xorm:"comment('播单图标未选中状态') VARCHAR(255)"`
	SowingIconSelected     string    `xorm:"comment('播单图标选中状态') VARCHAR(255)"`
	TvIconUnselected       string    `xorm:"comment('影视图标未选中状态') VARCHAR(255)"`
	TvIconSelected         string    `xorm:"comment('影视图标选中状态') VARCHAR(255)"`
	TreasureIconUnselected string    `xorm:"comment('宝箱图标未选中状态') VARCHAR(255)"`
	TreasureIconSelected   string    `xorm:"comment('宝箱图标选中状态') VARCHAR(255)"`
	MineIconUnselected     string    `xorm:"comment('我的图标未选中状态') VARCHAR(255)"`
	MineIconSelected       string    `xorm:"comment('我的图标选中状态') VARCHAR(255)"`
	Attr1                  string    `xorm:"VARCHAR(255)"`
	Attr2                  string    `xorm:"VARCHAR(255)"`
	Attr3                  string    `xorm:"VARCHAR(255)"`
	Attr4                  string    `xorm:"VARCHAR(255)"`
}

type KsMessage struct {
	Id         string    `xorm:"not null pk VARCHAR(255)"`
	MsgId      string    `xorm:"not null comment('消息推送id') VARCHAR(255)"`
	Title      string    `xorm:"comment('消息标题') VARCHAR(255)"`
	Intro      string    `xorm:"comment('消息简介') VARCHAR(255)"`
	Content    string    `xorm:"comment('内容链接 或者富文本') VARCHAR(255)"`
	CreateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	Status     int       `xorm:"not null comment('状态（已发：1，未发：0）') INT(11)"`
	Attr1      string    `xorm:"VARCHAR(255)"`
	Attr2      string    `xorm:"VARCHAR(255)"`
	Attr3      string    `xorm:"VARCHAR(255)"`
}

type LogCommentLike struct {
	Id         string    `xorm:"not null pk VARCHAR(255)"`
	CommentId  string    `xorm:"not null index VARCHAR(255)"`
	UserId     string    `xorm:"not null index VARCHAR(255)"`
	Status     int       `xorm:"not null default 1 comment('状态；-1已删除, 1正常') INT(11)"`
	CreateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

type Post struct {
	Id          string    `xorm:"not null pk VARCHAR(255)"`
	PostTypeId  string    `xorm:"not null index VARCHAR(255)"`
	UserId      string    `xorm:"not null index VARCHAR(255)"`
	Author      string    `xorm:"not null VARCHAR(255)"`
	Title       string    `xorm:"not null VARCHAR(255)"`
	Subtitle    string    `xorm:"not null VARCHAR(255)"`
	ImgUrl      string    `xorm:"VARCHAR(255)"`
	Context     string    `xorm:"not null TEXT"`
	Sort        int       `xorm:"not null default 0 comment('排序；值大靠前') INT(11)"`
	Choice      int       `xorm:"not null default 1 comment('精选（1：普通，2：精选）') INT(11)"`
	LookNumber  int       `xorm:"not null default 0 INT(11)"`
	Release     int       `xorm:"not null default 1 comment('是否发布；1：未发布，2：发布') INT(11)"`
	ReleaseDate time.Time `xorm:"DATETIME"`
	Status      int       `xorm:"not null default 1 comment('状态；-1已删除, 1正常') INT(11)"`
	CreateDate  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

type PostType struct {
	Id         string    `xorm:"not null pk VARCHAR(255)"`
	Name       string    `xorm:"not null VARCHAR(255)"`
	ImgUrl     string    `xorm:"VARCHAR(255)"`
	Sort       int       `xorm:"not null default 0 comment('排序；值大靠前') INT(11)"`
	Status     int       `xorm:"not null default 1 comment('状态；-1已删除, 1正常') INT(11)"`
	CreateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

type Resource struct {
	Id         string    `xorm:"not null pk VARCHAR(255)"`
	Url        string    `xorm:"not null VARCHAR(255)"`
	ForeignId  string    `xorm:"not null comment('外键，关联数据id') index VARCHAR(255)"`
	Type       string    `xorm:"comment('类型') index VARCHAR(255)"`
	Status     int       `xorm:"not null default 1 INT(11)"`
	CreateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

type Type struct {
	Id   string `xorm:"not null pk VARCHAR(255)"`
	Type string `xorm:"VARCHAR(255)"`
}

type Users struct {
	Id       string    `xorm:"not null pk VARCHAR(255)"`
	Name     string    `xorm:"VARCHAR(255)"`
	Age      int64     `xorm:"BIGINT(64)"`
	Birthday time.Time `xorm:"DATE"`
	GroupId  string    `xorm:"VARCHAR(255)"`
	TypeId   string    `xorm:"VARCHAR(255)"`
}
