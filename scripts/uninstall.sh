#!/usr/bin/bash

DIR="/usr/local/gomeasure"
EXEC="/usr/local/bin/gomeasure"

if [ "$(id -u)" != "0" ]; then
    # clear any previous sudo permission
    sudo -k
    echo "gomeasure will be permanently removed from your system."
    sudo echo # Echo to make user enter password
fi

sudo rm -rf "$DIR" "$EXEC"
echo "gomeasure successfully uninstalled"
