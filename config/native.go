package config

/**
 * @author eyesYeager
 * @date 2023/2/8 18:01
 */

type Native struct {
	Profiles Profiles `mapstructure:"profiles" json:"profiles" yaml:"profiles"`
	Nacos    Nacos    `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
}
