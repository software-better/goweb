docker stop -s goweb
docker rm -s goweb
docker rmi -s goweb
docker build -t goweb .
#docker run -it --rm -p 8300:8300 --name goweb goweb
echo Next step - run "(docker run -it --rm -p 8300:8300 --name goweb goweb)"


