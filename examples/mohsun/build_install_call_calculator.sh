#!/bin/sh
full_path=$(realpath $0)
#echo $full_path
 
dir_path=$(dirname $full_path)
#echo $dir_path
#istenilen yerden skripti chaghirmaq
cd $dir_path/calculator

#echo $HOME/go/bin
export PATH="$PATH:$HOME/go/bin"
echo $PATH | grep -o $HOME/go/bin
go mod tidy
echo "build edek"
go build .
echo "install edek"
go install #burda uje path var he

echo "calculatoru cahgiraq"
calculator

#exi sen go nu da rm eledin ee, hee yox yada)okdi