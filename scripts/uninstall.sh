if [ "$(id -u)" != "0" ]; then
    # clear any previous sudo permission
    sudo -k
    echo "gomeasure will be permanently removed from your system."
    sudo echo # Echo to make user enter password
fi

sudo rm -rf /usr/local/gomeasure /usr/bin/gomeasure 
echo "gomeasure successfully uninstalled"
