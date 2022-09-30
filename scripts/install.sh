#!/usr/bin/bash

set -e

DIR="/usr/local/gomeasure"
VERSION="0.1"

if [ "$(id -u)" != "0" ]; then
    # clear any previous sudo permission
    sudo -k
    echo "This script requires superuser access to install cmd tool."
    echo "You will be prompted for your password by sudo."
    sudo echo # Echo to make user enter password
fi
#  set -ex

# download zipped
curl "https://github.com/lordvidex/gomeasure/releases/download/v${VERSION}/gomeasure_${VERSION}_linux_amd64.tar.gz" -L -o zipped.tar.gz;

# move to lib directory

# clear folder and uninstall old versions if exists
[ -d "$DIR" ] && sudo rm -rf "$DIR"
[ -L "/usr/bin/gomeasure" ] && sudo rm /usr/bin/gomeasure

# create folder if it doesn't exist
[ ! -d "$DIR" ] && sudo mkdir -p "$DIR"

sudo tar -xf zipped.tar.gz -C "$DIR"
rm zipped.tar.gz # clean

# create symbolic links
sudo ln -s "$DIR/gomeasure" /usr/bin/gomeasure

# test that it works
echo "gomeasure successfully installed"
gomeasure --version
