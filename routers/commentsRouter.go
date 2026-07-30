package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:APISessionController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:APISessionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:APISessionController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:APISessionController"],
        beego.ControllerComments{
            Method: "Kill",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:APISignalController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:APISignalController"],
        beego.ControllerComments{
            Method: "Send",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:APISysloadController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:APISysloadController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/certificates`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/certificates`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Download",
            Router: `/certificates/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Burn",
            Router: `/certificates/burn/:key/:serial/:tfaname`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Restart",
            Router: `/certificates/restart`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Revoke",
            Router: `/certificates/revoke/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Renew",
            Router: `/certificates/revoke/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:DangerController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:DangerController"],
        beego.ControllerComments{
            Method: "RestartContainer",
            Router: `/container/restart`,
            AllowHTTPMethods: []string{"RestartContainer"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:DangerController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:DangerController"],
        beego.ControllerComments{
            Method: "DeletePKI",
            Router: `/pki/delete/:key`,
            AllowHTTPMethods: []string{"DeletePKI"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:DangerController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:DangerController"],
        beego.ControllerComments{
            Method: "InitPKI",
            Router: `/pki/init/:key`,
            AllowHTTPMethods: []string{"InitPKI"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:OVClientConfigController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:OVClientConfigController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/ov/clientconfig/edit`,
            AllowHTTPMethods: []string{"Edit"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:OVConfigController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:OVConfigController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/ov/config`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:OVConfigController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:OVConfigController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/ov/config`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:OVConfigController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:OVConfigController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/ov/config/edit`,
            AllowHTTPMethods: []string{"Edit"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/profile`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/profile/create`,
            AllowHTTPMethods: []string{"Create"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/profile/delete/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/glaolle/openvpn-ui/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "EditUser",
            Router: `/profile/edit/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
