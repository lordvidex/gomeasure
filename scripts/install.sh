#!/usr/bin/bash

set -e

DIR="/usr/local/gomeasure"
EXEC="/usr/local/bin/gomeasure"

if [ "$(id -u)" != "0" ]; then
  # clear any previous sudo permission
  sudo -k
  echo "This script requires superuser access to install cmd tool."
  echo "You will be prompted for your password by sudo."
  sudo echo # Echo to make user enter password
fi
#  set -ex

# download zipped
curl -s https://api.github.com/repos/lordvidex/gomeasure/releases/latest |
  grep "browser_download_url.*_linux_amd64.tar.gz" |
  cut -d : -f 2,3 |
  tr -d \" |
  wget -O zipped.tar.gz -qi -

# move to lib directory

# clear folder and uninstall old versions if exists
[ -d "$DIR" ] && sudo rm -rf "$DIR"
[ -L "$EXEC" ] && sudo rm "$EXEC"

# create folder if it doesn't exist
[ ! -d "$DIR" ] && sudo mkdir -p "$DIR"

sudo tar -xf zipped.tar.gz -C "$DIR"
sudo rm zipped.tar.gz # clean

# create symbolic links
sudo ln -s "$DIR/gomeasure" "$EXEC"

# test that it works
echo "gomeasure successfully installed"
gomeasure --version
