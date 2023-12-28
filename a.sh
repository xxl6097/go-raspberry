#!/bin/bash

function version() {
  version=$(cat version)
  if [ "$version" = "" ]; then
    version="0.0.0"
  fi
  v3=$(echo $version | awk -F'.' '{print($3);}')
  v2=$(echo $version | awk -F'.' '{print($2);}')
  v1=$(echo $version | awk -F'.' '{print($1);}')
  if [[ $(expr $v3 \>= 9) == 1 ]]; then
    v3=0
    if [[ $(expr $v2 \>= 9) == 1 ]]; then
      v2=0
      v1=$(expr $v1 + 1)
    else
      v2=$(expr $v2 + 1)
    fi
  else
    v3=$(expr $v3 + 1)
  fi
  echo "$v1.$v2.$v3" >version
}
version
