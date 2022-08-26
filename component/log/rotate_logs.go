package log

import (
	"github.com/chuangxinyuan/fengchu/config"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/cast"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct{}

func (r *fileRotateLogs) GetWriteSyncer(config *config.Zap, level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotateLogs.New(
		path.Join(config.Director, "%Y-%m-%d", level+".log"),
		rotateLogs.WithClock(rotateLogs.Local),
		rotateLogs.WithMaxAge(time.Duration(cast.ToInt(config.MaxAge))*24*time.Hour), // 日志留存时间
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
}
