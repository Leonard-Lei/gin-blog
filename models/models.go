package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
)

var db *gorm.DB

type Model struct {
	ID         int               `gorm:"primary_key" json:"id"`
	CreateTime util.JsonDateTime `json:"create_time" time_format:"2006-01-02 15:04:05"`
	UpdateTime util.JsonDateTime `json:"update_time" time_format:"2006-01-02 15:04:05"`
	DeleteFlag int               `json:"delete_flag"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:create_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

//update TimeStamp For Create Callback will set `CreateTime` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {

	nowTime := util.JsonDateTime(time.Now())

	err := scope.SetColumn("CreateTime", nowTime)
	if err != nil {
		log.Println(err)
	}
	err = scope.SetColumn("UpdateTime", util.JsonDateTime(time.Now()))
	if err != nil {
		log.Println(err)
	}
	log.Println("CreateTime callback")

	//检查是否有错误
	if !scope.HasError() {
		//nowTime := time.Now().Unix()

		//通过scope.FieldByName()获取所有字段，判断当前是否包含所需要字段
		// if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
		// 	//判断该字段的值是否为空
		// 	if createTimeField.IsBlank {
		// 		err := createTimeField.Set(nowTime)
		// 		if err != nil {
		// 			log.Println(err)
		// 		}
		// 	}
		// }

		// if modifyTimeField, ok := scope.FieldByName("UpdateTime"); ok {

		// 	scope.SetColumn("UpdateTime", util.JsonDateTime(time.Now()).String())
		// 	if modifyTimeField.IsBlank {
		// 		err := modifyTimeField.Set(nowTime)
		// 		if err != nil {
		// 			log.Println(err)
		// 		}
		// 	}
		// }
	}
}

// updateTimeStampForUpdateCallback will set `UpdateTime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//假设没有指定 update_column 的字段，我们默认在更新回调设置 ModifiedOn 的值
	//scope.SetColumn("UpdateTime", time.Now().Unix())
	scope.SetColumn("UpdateTime", util.JsonDateTime(time.Now()).String())
	scope.SetColumn("CreateeTime", util.JsonDateTime(time.Now()).String())
}

// deleteCallback will set `DeletedTime` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		//检查是否手动指定了delete_option
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		//获取我们约定的删除字段，若存在则 UPDATE 软删除，若不存在则 DELETE 硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeleteFlag")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				//返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				//添加值作为SQL的参数，也可用于防范SQL注入
				//scope.AddToVars(time.Now().Unix()),
				scope.AddToVars(util.JsonDateTime(time.Now()).String()),
				//返回组合好的条件SQL，看一下方法原型很明了
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
