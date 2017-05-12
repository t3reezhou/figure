#!/bin/bash
export TOP=$(pwd)
echo TOP:$TOP
export ROOT=${TOP/\/src\/github.com\/t3reezhou\/figure/}
echo ROOT:$ROOT
export GOBIN=$TOP/bin
echo GOBIN:$GOBIN
export GOPATH="$ROOT"
echo GOPATH:$GOPATH
