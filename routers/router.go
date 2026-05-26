// Package routers defines application routes
// @APIVersion 1.0.0
// @Title OpenVPN API
// @Description REST API allows you to control and monitor your OpenVPN server
// @Contact adam.walach@gmail.com
// License Apache 2.0
// LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/d3vilh/openvpn-ui/controllers"
	"github.com/d3vilh/openvpn-ui/state"
)

func Init(configDir string) {

	web.SetStaticPath("/swagger", "swagger")

	ns := web.NewNamespace(state.GlobalCfg.AutoPrefix,
		web.NSRouter("/", &controllers.MainController{}),
		web.NSRouter("/login", &controllers.LoginController{}, "get:Login;post:Login"),
		web.NSRouter("/logout", &controllers.LoginController{}, "get:Logout"),
		web.NSRouter("/auth/google", &controllers.LoginController{}, "get:GoogleLogin"),
		web.NSRouter("/auth/google/callback", &controllers.LoginController{}, "get:GoogleCallback"),
		web.NSRouter("/profile", &controllers.ProfileController{}),
		web.NSRouter("/settings", &controllers.SettingsController{}),
		web.NSRouter("/ov/config", &controllers.OVConfigController{ConfigDir: configDir}),
		web.NSRouter("/logs", &controllers.LogsController{}),
		web.NSRouter("/ov/clientconfig", &controllers.OVClientConfigController{ConfigDir: configDir}),
		web.NSRouter("/easyrsa/config", &controllers.EasyRSAConfigController{ConfigDir: configDir}),
		web.NSRouter("/dangerzone", &controllers.DangerController{}),

		web.NSInclude(&controllers.CertificatesController{ConfigDir: configDir}),

		web.NSNamespace("/api/v1",
			web.NSNamespace("/session", web.NSInclude(&controllers.APISessionController{})),
			web.NSNamespace("/sysload", web.NSInclude(&controllers.APISysloadController{})),
			web.NSNamespace("/signal", web.NSInclude(&controllers.APISignalController{})),
		),
	)
	web.AddNamespace(ns)
}
