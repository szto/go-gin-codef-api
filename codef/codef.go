package codef

import (
	config "go-gin-codef-api/config"

	ecg "github.com/codef-io/easycodefgo"
)

func GetCodef() *ecg.Codef {
	config := config.InitConfig()

	codef := &ecg.Codef{
		PublicKey: config.CodefPublicKey,
	}

	codef.SetClientInfoForDemo(config.CodefClientId, config.CodefClientSecret)

	return codef
}
