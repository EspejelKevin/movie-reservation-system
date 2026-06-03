package settings

import "github.com/alicanli1995/conform"

type Settings struct {
	AppPort string `conform:"env=APP_PORT,required"`
	DbUri   string `conform:"env=DB_URI,required"`
	Key     string `conform:"env=KEY,required"`
}

func NewSettings() (*Settings, error) {
	return conform.LoadGeneric[Settings](conform.FromEnv())
}
