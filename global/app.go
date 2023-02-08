package global

import (
	"eyesStars/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 16:53
 */

var Native = new(config.Native)

var Config = new(config.Config)

var DB = new(gorm.DB)

var Log = new(zap.Logger)
