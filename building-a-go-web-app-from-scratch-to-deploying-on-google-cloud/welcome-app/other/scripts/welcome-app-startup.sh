#! bin/bash

# Install Linux Dependencies
sudo apt-get update

# Install Go
wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz

#Install Git 
sudo apt install -y git-all

# Set the Export Path
export PATH=$PATH:/usr/local/go/bin

# pull from our repository
git clone https://github.com/martinomburajr/medium.git

#Change directory to our app
cd medium/building-a-go-web-app-from-scratch-to-deploying-on-google-cloud/welcome-app/

#Run the application
go run main.go


