#!/bin/bash
echo 'webhooks code automatic deployment'
case $1 in
    "refs/heads/master") # 可根据不同的分支做不同的任务
    echo $1
    ;;
esac