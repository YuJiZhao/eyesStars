package bootstrap

import (
	"eyesStars/app/model/entity"
	"eyesStars/global"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 17:23
 */

func InitializeDB() *gorm.DB {
	// 根据驱动配置进行初始化
	switch global.Config.Database.Driver {
	case "mysql":
		return initMySqlGorm()
	default:
		panic("默认不支持该数据库驱动，请自行拓展")
	}
}

// 初始化 mysql gorm.DB
func initMySqlGorm() *gorm.DB {
	dbConfig := global.Config.Database

	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外键约束
		Logger:                                   getGormLogger(), // 使用自定义 Logger
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	}); err != nil {
		global.Log.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		initMySqlTables(db)
		return db
	}
}

// 数据库表初始化
func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.Context{},
		entity.LogContext{},
		entity.LogFile{},
		entity.LogLogin{},
		entity.LogStar{},
		entity.LogVisit{},
		entity.Star{},
	)
	if err != nil {
		global.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.Config.Database.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.Config.Log.RootDir + "/" + global.Config.Database.LogFolder + "/" + time.Now().Format("2006-01-02") + ".log",
			MaxSize:    global.Config.Log.MaxSize,
			MaxBackups: global.Config.Log.MaxBackups,
			MaxAge:     global.Config.Log.MaxAge,
			Compress:   global.Config.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                      // 慢 SQL 阈值
		LogLevel:                  logMode,                                     // 日志级别
		IgnoreRecordNotFoundError: false,                                       // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.Config.Database.EnableFileLogWriter, // 禁用彩色打印
	})
}
