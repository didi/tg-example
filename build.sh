#!/bin/bash

# var
WORKSPACE=$(cd `dirname $0` && pwd -P)
DATE=$(date "+%s")
GOPATH=/tmp/go-build${DATE}

# 项目名称，根据具体项目改动
MODULE_NAME=tg-example
MODULE_PATH=${GOPATH}/src/github.com/didi

# env
export GOPATH
export GOROOT=/usr/local/go1.21.5
export PATH=${GOROOT}/bin:$GOPATH/bin:${PATH}:$GOBIN

# 添加go modules相关环境变量
export GOPROXY=http://goproxy.intra.xiaojukeji.com
export GOSUMDB=off
export GO111MODULE=auto

function build() {
    rm -rf $MODULE_PATH/$MODULE_NAME &> /dev/null
    mkdir -p $GOPATH/bin
    mkdir -p $MODULE_PATH
    ln -sf $WORKSPACE ${MODULE_PATH}/${MODULE_NAME}
    cd $MODULE_PATH/$MODULE_NAME
    go clean -modcache

    echo "Building……" && make

    if [[ $? != 0 ]];then
        echo -e "Build failed !"
        exit 1
    fi
    echo -e "Build success!"
}

build
