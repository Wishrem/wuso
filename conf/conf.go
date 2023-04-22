package conf

import (
	"bytes"
	_ "embed"
	"flag"
	"io"
	"os"
	"path/filepath"
	"strings"

	conf "github.com/Wishrem/wuso/conf/type"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	flagconf *string
	config   = new(conf.Config)

	// go:embed dev_config.toml
	devConfigBytes []byte
	// go:embed pro_config.toml
	proConfigBytes []byte
)

func init() {
	flagconf = flag.String("conf", "", "select config file:\n dev:debug\n pro:release\n")
	flag.Parse()

	var r io.Reader
	s := read(&r)

	viper.SetConfigType("toml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	save(s)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func read(r *io.Reader) string {
	switch strings.ToLower(strings.TrimSpace(*flagconf)) {
	case "dev":
		*r = bytes.NewReader(devConfigBytes)
		return "dev"
	case "pro":
		*r = bytes.NewReader(proConfigBytes)
		return "pro"
	default:
		*r = bytes.NewReader(devConfigBytes)
		return "dev"
	}
}

func save(s string) {
	viper.SetConfigName(s + "_config")
	viper.AddConfigPath("./conf")
	file := "./conf/" + s + "_config.toml"

	_, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(file), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(file)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}
}

func Get() conf.Config {
	return *config
}
