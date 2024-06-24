#!/bin/bash
source .env
echo -n "Please number of the day: "
read number

dir=Day_$number
mkdir $dir

cd $dir

touch main.go
echo "package main" > main.go
echo "
func main() {

}" >> main.go
touch example
curl --cookie "session=$AOC_COOKIE" https://adventofcode.com/2023/day/$number/input > input
go mod init $dir

cd ..
go work use $dir

