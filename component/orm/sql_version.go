package sql

import (
	"errors"
	"fengchu/utils"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"runtime"
	"strings"
	"time"
)

// Data .
type Data struct {
	db *gorm.DB
}

func (FlywaySchemaHistory) TableName() string {
	return "flyway_schema_history"
}

type FlywaySchemaHistory struct {
	InstalledRank int64     `gorm:"primary_key;not null;comment:'ID';auto_increment;" db:"installed_rank" json:"installed_rank" form:"installed_rank"`
	Version       string    `gorm:"column:version" db:"version" json:"version" form:"version"`
	Description   string    `gorm:"column:description" db:"description" json:"description" form:"description"`
	Type          string    `gorm:"column:type" db:"type" json:"type" form:"type"`
	Script        string    `gorm:"column:script" db:"script" json:"script" form:"script"`
	Checksum      string    `gorm:"column:checksum" db:"checksum" json:"checksum" form:"checksum"`
	InstalledBy   string    `gorm:"column:installed_by" db:"installed_by" json:"installed_by" form:"installed_by"`
	InstalledOn   time.Time `gorm:"column:installed_on" db:"installed_on" json:"installed_on" form:"installed_on"`
	ExecutionTime int64     `gorm:"column:execution_time" db:"execution_time" json:"execution_time" form:"execution_time"`
	Success       int64     `gorm:"column:success" db:"success" json:"success" form:"success"`
}

func createFlywayTable(data *Data) {
	err := data.db.Migrator().CreateTable(&FlywaySchemaHistory{})
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		log.Error("Failed to auto migrate table. Error: ", err.Error())
	} else {
		log.Info("Successfully created table flyway_schema_history ")
	}
}

// Flyway SQL文件版本管理
func Flyway(data *Data) {
	createFlywayTable(data)
	var dir string
	if runtime.GOOS == "windows" {
		dir = "../../docs/sql/"
	} else {
		dir = "docs/sql/"
	}
	names := utils.GetSortedSQLFiles(dir)
	for _, name := range names {
		fullName := dir + name
		checkSum := utils.GetFileMD5(fullName)
		sqls, readFileErr := ioutil.ReadFile(fullName)
		if readFileErr != nil {
			panic("failed to read SQL file. Error: " + readFileErr.Error())
		}
		count, err := checkFileUpdate(data, name, checkSum)
		if count == 0 && err == nil {
			sqlArr := strings.Split(string(sqls), ";")
			for _, sql := range sqlArr {
				sql = strings.TrimSpace(sql)
				if sql == "" {
					continue
				}
				err := data.db.Exec(sql).Error
				if err != nil {
					panic("failed to execute SQL. Error: " + err.Error())
				}
			}
			flywayHistory := FlywaySchemaHistory{
				Version:       utils.GetVersion(name),
				Description:   utils.GetDescription(name),
				Type:          "SQL",
				Script:        name,
				Checksum:      checkSum,
				InstalledBy:   "root",
				InstalledOn:   time.Now(),
				ExecutionTime: time.Now().Unix(),
				Success:       1,
			}
			e := updateHistory(data, flywayHistory)
			if e != nil {
				panic("failed to update history. Error: " + e.Error())
			}
			log.Info("executed SQL file " + name + " successfully")
		} else if err != nil {
			panic("failed to execute SQL. Error: " + err.Error())
		}
	}
}

func updateHistory(data *Data, flywayHistory FlywaySchemaHistory) error {
	err := data.db.Create(&flywayHistory).Error
	return err
}

// 检查SQL文件是否有更新。如果有更新就报错
func checkFileUpdate(data *Data, name string, checkSum string) (int64, error) {
	var flywayHistory FlywaySchemaHistory
	var count int64
	err := data.db.Where("script = ?", name).Find(&flywayHistory).Limit(1).Count(&count).Error
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, nil
	} else {
		if flywayHistory.Checksum != checkSum {
			return 1, errors.New("SQL file " + name + " already been updated.")
		} else {
			return 1, nil
		}
	}
}
