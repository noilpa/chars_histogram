#!/bin/bash

CYAN_COLOR='\033[0;36m'
NO_COLOR='\033[0m'

echo -e "${CYAN_COLOR}===> Generate directories and files <===${NO_COLOR}"
python DirsAndFilesGenerator.py 20
echo -e "${CYAN_COLOR}===> Frequence histogram <===${NO_COLOR}"
go run src/cmd/main.go $(pwd)/dir_1
echo -e "${CYAN_COLOR}===> Remove directories and files <===${NO_COLOR}"
rm -rf dir_1
