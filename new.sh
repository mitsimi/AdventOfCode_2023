#!/bin/bash

echo -n "Please number of the day: "
read number

dir=Day_$number
mkdir $dir

cd $dir

touch main.go
echo "package main" > main.go
touch example
touch input
go mod init $dir

cd ..
go work use $dir

