#! bin/bash

# Install Linux Dependencies
sudo apt-get update && sudo apt-get upgrade

# Install Go
wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz

# Set the Export Path
export PATH=$PATH:/usr/local/go/bin

# pull from our repository


