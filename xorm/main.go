package main

import (
	"fmt"
	"gin/log"
	"gin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

//engine 是goroutine安全的
var engine *xorm.Engine
var engines []*xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@(localhost:3306)/hlj?charset=utf8")
	_ = engine.Ping()
	err = engine.Sync2(new(models.Users), new(models.Group), new(models.Type))
	if err != nil {
		logrus.Info(err)
	}
	//日志
	//控制台显示sql  默认INFO
	engine.ShowSQL(true)
	//控制台打印调试及以上的信息
	engine.Logger().SetLevel(core.LOG_INFO)
	//写入文件
	f, err := os.OpenFile("api.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logrus.Info(err.Error())
		return
	}
	mw := io.MultiWriter(os.Stdout, f)
	engine.SetLogger(xorm.NewSimpleLogger(mw))
	//日志记录到syslog
	//logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	//if err != nil {
	//	log.Fatalf("Fail to create xorm system logger: %v\n", err)
	//}
	//
	//logger := xorm.NewSimpleLogger(logWriter)
	//logger.ShowSQL(true)
	//engine.SetLogger(logger)

	//连接池
	engine.SetMaxIdleConns(100) //连接池的空闲数大小
	engine.SetMaxOpenConns(300) //最大打开连接数
	//engine.SetConnMaxLifetime()

	//名称映射规则
	//core.SameMapper和core.GonicMapper。 * SnakeMapper 支持struct为驼峰式命名，
	// 表结构为下划线命名之间的转换，这个是默认的Maper；
	// * SameMapper 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名；
	// * GonicMapper 和SnakeMapper很类似，但是对于特定词支持更好，比如ID会翻译成id而不是i_d。
	//engine.SetMapper(core.SameMapper{})
	//表名称和字段名称的映射规则默认是相同的，当然也可以设置为不同，如：
	engine.SetTableMapper(core.SnakeMapper{})
	engine.SetColumnMapper(core.GonicMapper{})
	//如果希望所有的表名都在结构体自动命名的基础上加一个前缀而字段名不加前缀，
	// 则可以在engine创建完成后执行以下语句：
	//tbMapper:=core.NewPrefixMapper(core.SnakeMapper{},"prefix_")
	//engine.SetTableMapper(tbMapper)
	//类似的还有后缀、缓存映射 core.NewSuffixMapper() core.NewCacheMapper()

	// 使用Table和Tag改变名称映射
	//如果结构体拥有TableName() string的成员方法，那么此方法的返回值即是该结构体对应的数据库表名。
	//
	//通过engine.Table()方法可以改变struct对应的数据库表的名称，
	// 通过sturct中field对应的Tag中使用xorm:"'column_name'"可以使该field对应的Column名称为指定名称。
	// 这里使用两个单引号将Column名称括起来是为了防止名称冲突，
	// 因为我们在Tag中还可以对这个Column进行更多的定义。如果名称不冲突的情况，
	// 单引号也可以不使用。

	//表名的优先级顺序如下：
	//
	//engine.Table() 指定的临时表名优先级最高
	//TableName() string 其次
	//Mapper 自动映射的表名优先级最后
	//字段名的优先级顺序如下：
	//
	//结构体tag指定的字段名优先级较高
	//Mapper 自动映射的表名优先级较低

	//log := logrus.New()
	//f1, _ := os.OpenFile("api.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	//log.SetOutput(io.MultiWriter(os.Stdout, f1))
	log.Info(engine.DBMetas())
	//log.Info(engine.TableInfo(&Users{}))

	//CreateIndexes CreateUniques创建索引和唯一索引
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	log.Info(engine.Tables)
	var ksHelp []models.KsHelp
	//bool, err := engine.ID("0bc4ad4b-d9c0-475e-a630-624b83009da1").Get(&ksHelp)
	//log.Info(err, bool)
	//log.Info(ksHelp)
	engine.Logger().ShowSQL(true)
	fmt.Println(engine.Logger().IsShowSQL())
	_ = engine.Alias("o").Where("o.title=?", "de毛东0").Desc("create_date").Find(&ksHelp)
	log.Info(ksHelp)
	var ksFeedbacks []models.KsFeedback
	//执行指定的Sql语句，并把结果映射到结构体。有时，当选择内容或者条件比较复杂时，可以直接使用Sql
	_ = engine.SQL("select * from ks_feedback").Find(&ksFeedbacks)
	log.Info(ksFeedbacks)
	//engine.Find()
	var ksIconThemes []models.KsIconTheme
	_ = engine.In("status", -1, 0).Find(&ksIconThemes)
	log.Info(ksIconThemes)

	var ksIconTheme models.KsIconTheme
	_, _ = engine.Cols("status", "create_time").Get(&ksIconTheme)
	log.Info(ksIconTheme)
	//查询或更新所有字段，一般与Update配合使用，因为默认Update只更新非0，非”“，非bool的字段。
	ksIconTheme2 := models.KsIconTheme{ThemeTitle: "fsdssssf", Status: 0}
	//_, err = engine.AllCols().ID("5b93a472-6c1f-4405-bed7-2efe2e9849bf").Update(&ksIconTheme2)
	//log.Info(ksIconTheme2)
	_, _ = engine.Omit("status").Where("id=?", "5b93a472-6c1f-4405-bed7-2efe2e9849bf").Update(&ksIconTheme2)
	_, err = engine.Insert(&ksIconTheme2)
	log.Info(err)

	var ksIconThemes2 []models.KsIconTheme
	//engine.Distinct("theme_title").Find(&ksIconThemes2)
	_ = engine.Limit(2, 0).Find(&ksIconThemes2)
	//GroupBy(string) Having(string)
	log.Info(ksIconThemes2)

	//在当前语句有效
	engine.NoAutoTime() //created updated 字段不会赋值为当前时间
	engine.NoCache()    //非缓存模式执行
	engine.NoAutoCondition()
	engine.UseBool()
	//当从一个struct来生成查询条件或更新字段时，xorm会判断struct的field是否为0,“”,nil，
	// 如果为以上则不当做查询条件或者更新内容。因为bool类型只有true和false两种值，
	// 因此默认所有bool类型不会作为查询条件或者更新字段。如果可以使用此方法，如果默认不传参数，
	// 则所有的bool字段都将会被使用，如果参数不为空，
	// 则参数中指定的为字段名，则这些字段对应的bool值将被使用。
	engine.NoCascade() //是否自动关联查询field中的数据，如果struct的field也是一个struct并且映射为某个Id，
	// 则可以在查询时自动调用Get方法查询出对应的数据。
	//theme := new(models.KsIconTheme)

	//Exist 相较 Get 性能更好
	//如果你的需求是：判断某条记录是否存在，若存在，则返回这条记录。
	//
	//建议直接使用Get方法。
	//
	//如果仅仅判断某条记录是否存在，则使用Exist方法，Exist的执行效率要比Get更高。

	users := make([]UserGroup, 0)
	engine.Join("INNER", "`group`", "group.id=users.group_id").Find(&users)
	log.Info(users)
	users1 := make([]UserGroup, 0)
	engine.SQL("select users.*, g.name group_name from users, `group` g where users.group_id = g.id").Find(&users1)
	log.Info(users1)
	//users2 := UserGroupType{}
	//engine.Join("INNER", "`group`", "group.id=users.group_id").
	//	Join("INNER", "type", "type.id=users.type_id").Find(&users2)
	//log.Info(users2)
	err = engine.Where("age > ? or name =?", 20, "xlw").Iterate(new(models.Users), func(i int, bean interface{}) error {
		user := bean.(*models.Users)
		log.Info(user)
		return err
	})

	users3 := models.Users{}
	count, err := engine.Count(&users3)
	fields := make(map[string]interface{}, 0)
	fields["count"] = count
	log.InfoWithFields("table users :count", fields)

	user := new(models.Users)
	rows, err := engine.Where("age>?", 13).Rows(user)
	if err != nil {

	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(user)
		userMap := make(map[string]interface{}, 0)
		userMap["user"] = user
		log.InfoWithFields("", userMap)
	}
	//var user2 models.Users

	//事务处理
	session := engine.NewSession()
	defer session.Close()
	err = session.Begin()
	user1 := models.Users{Name: "def", Birthday: time.Now(), Age: 23}
	_, err = session.Insert(&user1)
	if err != nil {
		session.Rollback()
		return
	}
	user2 := models.Users{Name: "xlw"}
	session.Exec("delete from users where name=?", user2.Name)
	if err != nil {
		session.Rollback()
		return
	}
	err = session.Commit()
	if err != nil {
		return
	}

	//缓存

}

//结构体中extends标记对应的结构顺序应和最终生成SQL中对应的表出现的顺序相同。
//对于不重复字段，可以{{.GroupId}}，对于重复字段{{.User.Id}}和{{.Group.Id}}
type UserGroup struct {
	models.Users `xorm:"extends"`
	models.Group `xorm:"extends"`
}

func (UserGroup) TableName() string {
	return "users"
}

type UserGroupType struct {
	models.Users `xorm:"extends"`
	models.Group `xorm:"extends"`
	models.Type  `xorm:"extends"`
}

func (UserGroupType) TableName() string {
	return "users"
}
