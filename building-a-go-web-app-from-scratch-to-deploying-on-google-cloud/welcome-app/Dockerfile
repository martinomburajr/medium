FROM "golang:alpine"
# Here the FROM clause states which base image we are intending to work with. If the image does not exist, locally, Docker automatically fetches it from Dockerhub. If you supplied a URI for the image, Docker will download it from there too. Here we begin with the xxx image

MAINTAINER "Martin Ombura <info@martinomburajr.com"
# Shows who created/maintains the file

CMD apt-get update
# CMD let's us run commands within the container once dowloaded

WORKDIR /go/src
# Tells Docker to create a working directory that the container will by default use for your project. When you docker -ti <image> it will check into this folder first

COPY /home/<your-username>/go/src/welcome-app /home/src
# This command tells Docker to copy files from our local machine, into the container that is being built. In some cases we can choose to download our code from Github or any other source, for our case it is simple enough to just COPY the welcome-app to Docker. and place it in the WORKDIR mentioned earlier.

EXPOSE 8080
# This tells Docker to expose a certain port that can be listened to. This is important since our application is exposing on Port 8080, we need Docker to also expose on this port so that external sources can interact with our app.

ENTRYPOINT "go" "run" "main.go"
# This is the first command to run once the container starts, turning it into an automatically running welcome-app server.





#To run use the command below
#docker build -t imagename:version .