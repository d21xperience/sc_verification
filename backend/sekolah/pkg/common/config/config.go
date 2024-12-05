package config

import (
	"sekolah/helper"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl      string `mapstructure:"DB_URL"`
	DBPort     uint16 `mapstructure:"PORT"`
	DBName     string `mapstructure:"DB_USER"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_USER"`
}

func LoadConfig() (*Config, error) {
	/*	tentukan posisi folder dimana file env tersimpan.
		Dalam Golang, tanda titik (.) pada penulisan path sering digunakan untuk merujuk ke direktori saat ini atau ke elemen yang berada di dalam direktori saat ini.
		Pada projek ini file env disimpan pada folder pkg/common/envs. Karena file main-nya berada pada folder cmd/main.go maka penulisan path-nya yaitu : "../pkg/common/envs".
		.. (dua titik) artinya merujuk ke direktori induk (scproject) dari direktori saat ini.
	*/
	viper.AddConfigPath("../pkg/common/envs")
	// Tentukan nama file beserta file extension-nya, jika tidak ada tulis extentions-nya diawali dengan tanda titik (.)
	viper.SetConfigName(".env")
	// Tuliskan file extentions-nya
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	helper.ErrorLog("Error pada viper\v", err)
	var c Config
	errs := viper.Unmarshal(&c)
	helper.ErrorPanic(errs)
	return &c, nil

}
