#!/bin/bash
# Exit immediately if a command exits with a non-zero status
set -e

# Directory where OpenVPN configuration files are stored
OPENVPN_DIR=$(grep -E "^OpenVpnPath\s*=" /opt/openvpn-ui/conf/app.conf | cut -d= -f2 | tr -d '"' | tr -d '[:space:]')
echo "Init. OVPN path: $OPENVPN_DIR"
EASY_RSA=$(grep -E "^EasyRsaPath\s*=" /opt/openvpn-ui/conf/app.conf | cut -d= -f2 | tr -d '"' | tr -d '[:space:]')

# Change to the OpenVPN GUI directory
cd /opt/openvpn-ui

# If the provisioned file does not exist in the OpenVPN directory, prepare the certificates and create the provisioned file
if [ ! -f $OPENVPN_DIR/.provisioned ]; then
  echo "Preparing certificates"
  make-cadir $EASY_RSA
  $EASY_RSA/easyrsa --pki-dir=$EASY_RSA/pki init-pki
  ln -s $EASY_RSA/pki $OPENVPN_DIR/pki
  mkdir -p $OPENVPN_DIR/ccd
  mkdir -p $OPENVPN_DIR/config
  mkdir -p $OPENVPN_DIR/clients

  echo 'set_var EasyRSADN	"org"
set_var EasyRSAReqCountry	"UA"
set_var EasyRSAReqProvince	"KY"
set_var EasyRSAReqCity		"Kyiv"
set_var EasyRSAReqOrg		"SweetHome"
set_var EasyRSAReqEmail		"sweet@home.net"
set_var EasyRSAReqOu		"MyOrganizationalUnit"
set_var EasyRSAReqCn		"server"
set_var EasyRSAKeySize		2048
set_var EasyRSACaExpire		3650
set_var EasyRSACertExpire	825
set_var EasyRSACertRenew	30
set_var EasyRSACrlDays		180' > "$OPENVPN_DIR/config/easy-rsa.vars"

  # Uncomment line below to generate CA and server certificates (should be done on the side of OpenVPN container or server however)
  ./scripts/generate_ca_and_server_certs.sh init_all

  # Create the provisioned file
  touch $OPENVPN_DIR/.provisioned
  echo "First OpenVPN UI start."
fi

# Create the database directory if it does not exist
mkdir -p db

# Start the OpenVPN GUI
echo "Starting OpenVPN UI!"
./openvpn-ui
