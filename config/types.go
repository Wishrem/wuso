package config

import "time"

type mysql struct {
	Host     string
	Port     string
	Database string
	UserName string `yaml:"username"`
	Password string
	Charset  string
}

type jwt struct {
	Secret     []byte `mapstructure:"-"`
	Issuer     string
	ExpireTime time.Duration `mapstructure:"expire-time"`
}

type config struct {
	MySQL  mysql  `yaml:"mysql"`
	JWT    jwt    `yaml:"jwt"`
	Server server `yaml:"server"`
}

type server struct {
	Addr           string
	ReadTimeout    time.Duration `yaml:"read-timeout"`
	WriteTimeout   time.Duration `yaml:"write-timeout"`
	MaxHeaderBytes int           `yaml:"max-header-bytes"`
	WithoutClient  bool          `yaml:"without-client"`
}
