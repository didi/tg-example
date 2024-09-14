#!/bin/bash
#############################################
## main
## 以托管方式, 启动服务
## control.sh脚本, 必须实现start方法
#############################################
workspace=$(cd $(dirname $0) && pwd -P)
cd $workspace
module=tg-example
app=$module

function choose_conf_file() {
    local clusterfile="$workspace/.deploy/service.cluster.txt"
    if [ -f $clusterfile ]; then
        local cluster=`cat $clusterfile`
        if [ $cluster == "hna-v" ]; then
            echo "Internal hna-v environment"
            cp $workspace/conf/hna-v/* $workspace/conf/
            return
        elif [ $cluster == "hna-pre-v" ]; then
            echo "Internal hna-pre-v environment"
            cp $workspace/conf/hna-pre-v/* $workspace/conf/
            return
        elif [ $cluster == "hnb-v" ]; then
            echo "Internal hnb-v production environment"
            cp $workspace/conf/hnb-v/* $workspace/conf/
            return
        elif [ $cluster == "hne-v" ]; then
            echo "Internal hne-v production environment"
            cp $workspace/conf/hne-v/* $workspace/conf/
            return
        elif [ $cluster == "test" ]; then
            echo "Test environment"
            cp $workspace/conf/test/* $workspace/conf/
            return
        else
           echo "no correct environment " $cluster
           exit 1
        fi
    else
        echo "no service cluster file"
        exit 1
    fi
}
function run_command() {
	action=$1
	case $action in
	    "start" )
        choose_conf_file
		# 启动服务, 以前台方式启动, 否则无法托管
		exec &> >(while read line; do echo "[$(date "+%Y-%m-%d %H:%M:%S")] $line"; done) ./bin/$app
		;;
	    * )
		# 非法命令, 已非0码退出
		echo "unknown command"
		exit 1
		;;
	esac
}

run_command $1

