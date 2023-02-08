package config

/**
 * @author eyesYeager
 * @date 2023/1/28 17:12
 */

type Log struct {
	Level           string
	RootDir         string
	LogFolderApp    string
	LogFolderServer string
	Format          string
	ShowLine        bool
	MaxBackups      int
	MaxSize         int
	MaxAge          int
	Compress        bool
}
