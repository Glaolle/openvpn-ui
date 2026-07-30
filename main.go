package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/glaolle/openvpn-ui/lib"
	"github.com/glaolle/openvpn-ui/models"
	"github.com/glaolle/openvpn-ui/routers"
	"github.com/glaolle/openvpn-ui/state"
)

func main() {
	configDir := flag.String("config", "conf", "Path to config dir")
	flag.Parse()

	configFile := filepath.Join(*configDir, "app.conf")
	logs.Info("Config file:", configFile)

	if err := web.LoadAppConfig("ini", configFile); err != nil {
		panic(err)
	}

	models.InitDB()
	models.CreateDefaultUsers()
	defaultSettings, err := models.CreateDefaultSettings()
	if err != nil {
		panic(err)
	}

	models.CreateDefaultOVConfig(*configDir, defaultSettings.OVConfigPath, defaultSettings.MIAddress, defaultSettings.MINetwork)
	models.CreateDefaultOVClientConfig(*configDir, defaultSettings.OVConfigPath, defaultSettings.MIAddress, defaultSettings.MINetwork)
	models.CreateDefaultEasyRSAConfig(*configDir, defaultSettings.EasyRSAPath, defaultSettings.MIAddress, defaultSettings.MINetwork)
	state.GlobalCfg = *defaultSettings

	pkiPath := filepath.Join(defaultSettings.OVConfigPath, "pki")
	defaultPKI, err := lib.CreateDefaultPKI(pkiPath)
	if err != nil {
		panic(err)
	}
	state.GlobalPKI = *defaultPKI

	_, err = os.Stat(filepath.Join(pkiPath, "ta.key"))
	if err != nil {
		if os.IsNotExist(err) {
			logs.Info("ta.key not exist!")
			lib.GenerateTA(filepath.Join(pkiPath, "ta.key"))
			lib.GenerateCA()
			lib.GenerateServerCert()
			lib.GenerateCRL()
		}
	}

	routers.Init(*configDir)

	lib.AddFuncMaps()
	web.Run()
}
