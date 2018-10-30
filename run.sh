#!/bin/bash

process=`ps aux|grep beego-demo|grep webapps|wc -l`

if [ $process -eq 1 ]
then
        pid=`ps aux|grep beego-demo|grep webapps|awk -F " " '{print $2}'`
        kill -9 $pid
        echo "killed old process"
else
        echo "no old process"
fi

echo "start new process"
nohup /webapps/myApp/beego-demo/beego-demo > /tmp/beego.log 2>&1 &


