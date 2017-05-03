#!/bin/bash

export QINGTOP=$(pwd)
export QINGROOT=${QINGTOP/\/src\/github.com\/t3reezhou\/figure/}
echo $QINGROOT
export GOBIN=$QINGTOP/bin
export GOPATH="$QINGROOT"
