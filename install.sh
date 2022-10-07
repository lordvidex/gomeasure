#!/usr/bin

{
    set -e

    readonly GITHUB_USERNAME=lordvidex
    readonly GITHUB_REPOSITORY=gomeasure

    if [ "$(id -u)" != "0" ]; then
      # clear any previous sudo permission
      sudo -k
      echo "This script requires superuser access to install apt packages."
      echo "You will be prompted for your password by sudo."
      sudo echo # Echo to make user enter password
    fi
  set -ex

  # Curl the public key and add it to apt-key locally
  curl -s --compressed "https://${GITHUB_USERNAME}.github.io/${GITHUB_REPOSITORY}/assets/KEY.gpg" | apt-key add -

  # Curls the path to `.deb` file and cat it to the apt `source.list`
  curl -s --compressed -o /etc/apt/sources.list.d/gomeasure.list "https://${GITHUB_USERNAME}.github.io/${GITHUB_REPOSITORY}/assets/gomeasure.list"

  # Update apt to get latest version of the gomeasure app
  apt update
  apt install gomeasure -y

  # Test cli
  echo "gomeasure installed to $(which gomeasure)}"
  #gomeasure --version
}
