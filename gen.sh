#!/bin/bash
mkdir $1
cp templates/README $1/README.md
cp templates/main $1/main.go
cd $1
touch main.go
