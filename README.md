# Guide

## Build image docker 
```
docker build -t ahfrd/kredit-plus:v1 .
```

## Run kredit-plus image on docker
```
docker run -d -p 9018:9018 -v config:/app/config --name kredit-plus-v1 ahfrd/kredit-plus:v1
```

### Optional
```
docker run -d -p 3306:3306 -p 33060:33060 -e MYSQL_ROOT_PASSWORD=root --name db mysql:8.0
```
