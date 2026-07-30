package state

import (
	"github.com/glaolle/openvpn-ui/models"
	"github.com/kemsta/go-easyrsa/v2/pki"
)

var GlobalCfg models.Settings
var GlobalPKI pki.PKI
