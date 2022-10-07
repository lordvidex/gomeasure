#!/usr/bin/bash

DIR="/usr/local/gomeasure"
EXEC="/usr/local/bin/gomeasure"

echo "gomeasure will be permanently removed from your system."

if [ "$(id -u)" != "0" ]; then
  sudo echo # echo to make user enter password
fi
sudo rm -rf "$DIR" "$EXEC"
echo "gomeasure successfully uninstalled"
