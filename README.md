# goweb - a simple containerized go application example

This simple program is meant to be used to demonstrate packaging a Go web application
using minimally sized runtime containers.

## Building

The build is taken care of by naive Dockerfile and build-run.sh bash script.
I would like to mimic the approach used in [go-image-test](https://gitlab.com/sirile/go-image-test) but the 
centurylink/golang-builder step creates a container that does not run properly.

~~~bash
#build-run.sh
docker build -t goweb .
~~~

## Running

After that running the container can be done with

~~~bash
docker run -it --rm -p 8300:8300 --name goweb goweb
~~~

## Building with golang-builder (to be steps):

For building the container
[golang-builder](https://github.com/CenturyLinkLabs/golang-builder) is used. The
container can be built (assuming you have a working Docker environment) with:

~~~bash
docker run --rm -v "$(pwd):/src" -v /var/run/docker.sock:/var/run/docker.sock centurylink/golang-builder
~~~

If you want to minimize and tag the image straight after the build the command
is

~~~bash
docker run --rm -v "$(pwd):/src" -v /var/run/docker.sock:/var/run/docker.sock -e COMPRESS_BINARY=true centurylink/golang-builder software-better/goweb 
~~~
