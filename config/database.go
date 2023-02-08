package config

/**
 * @author eyesYeager
 * @date 2023/1/28 17:21
 */

type Database struct {
	Driver              string
	Host                string
	Port                int
	Database            string
	UserName            string
	Password            string
	Charset             string
	MaxIdleConns        int
	MaxOpenConns        int
	LogMode             string
	EnableFileLogWriter bool
	LogFolder           string
}
