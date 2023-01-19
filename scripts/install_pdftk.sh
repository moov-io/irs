#!/bin/sh

OS_NAME=$TRAVIS_OS_NAME
if [ "$OS_NAME" = "" ]; then
  if [ $(uname -s) = "Darwin" ]; then
	  export OS_NAME=osx
  else
	  export OS_NAME=linux
  fi
fi

if [ "$OS_NAME" = "linux" ]; then
  echo 'pdftk install on linux'
  sudo apt update
  sudo apt install snapd
  sudo snap install pdftk
fi


if [ "$OS_NAME" = "osx" ]; then
  echo 'pdftk install on osx'
  wget -O pdftk_server-setup.pkg https://www.pdflabs.com/tools/pdftk-the-pdf-toolkit/pdftk_server-2.02-mac_osx-10.11-setup.pkg
  sudo installer -pkg pdftk_server-setup.pkg -tgt /
  rm -f pdftk_server-setup.pkg
fi
