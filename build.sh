#!/bin/bash
#这里可替换为你自己的执行程序，其他代码无需更改
APP_NAME=taozhugong

rm -rf output
mkdir -p output
go build -o "${APP_NAME}" main.go
cp run.sh output/
cp -rf conf output/
mv ${APP_NAME} output/
