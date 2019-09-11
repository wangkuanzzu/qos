#!/bin/bash
#脚本使用说明：
#1.脚本执行的环境：在我们要导出数据的节点服务器上，正在运行qosd程序，是否为验证节点无影响。
#2.脚本执行的参数：目标服务器ip。‘192.168.1.200,192.168.1.202’
#3.脚本功能描述：停止当前服务器qosd程序，导出停止时候的data数据；睡眠30s：模拟其他节点依旧正常打块；停止目标节点qosd程序，将data数据传输给目标节点并覆盖。

if [ ! -n "$1" ]; then
    echo "need ips for deliver data"
    exit 1
fi

IFS="," HOSTS=($1)
for h in ${HOSTS[@]};do
  echo "***目标节点:$h"
done

ip=$(ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6 | awk '{print $2}')
#echo "<<<------------本机IP--------------->>>$ip"

#--------------------------------------------------------------------------------

QOSCLI="~/gocode/bin/qoscli"
QOSD="~/gocode/bin/qosd"
QOSPWD="11111111"

SSHCOMMAND="sshpass -pvagrant ssh vagrant@HOST \"COMMAND\""

#run_command host command
#ex. run_command 192.168.1.224 "$QOSCLI q validators --indent"
run_command(){
  c=${SSHCOMMAND/HOST/$1}
  c=${c/COMMAND/$2}
  eval $c
}

random(){
  ((r=$RANDOM%$1))
  echo $r
}

#------------------------------------------------------------------------------------

#1. 停止192.168.1.201的qosd进程，记录当时的高度为H。导出该节点的所有data数据：不包含priv_validator_state.json
echo "***操作当前服务器***"
pkill qosd
rm -rf ~/data*

cp -r ~/.qosd/data ~/data_H
rm -rf  ~/data_H/priv_validator_state.json

#2. 将data数据进行压缩:data.tar 
tar -cf  ~/data.tar ~/data_H

echo "***睡眠30s***"
sleep 30s

echo "***循环操作目的地址***"
#循环目的地址
for h in ${HOSTS[@]};do
  echo $h

  run_command $h "rm -rf ~/data*"
  #3. 将data.tar进行传输给ip：192.168.1.200
  sshpass -pvagrant scp ~/data.tar vagrant@$h:/home/vagrant/data.tar

  #4. IP为192.168.1.200的节点拥有全网2/3以上的votingpower，故其依然正常打块。其数据比data.tar中要多。

  #5. 由于我们决定恢复至192.168.1.201停止时的高度。
	#所以停止192.168.1.200的qosd进程，
	#执行命令qosd unsafe-reset-all清除201上的所有数据，
	#将data.tar中的数据替换./qosd/data/下的数据。
  run_command $h "pkill qosd"
  run_command $h "~/gocode/bin/qosd unsafe-reset-all"
  run_command $h "tar -xf ~/data.tar -C ~/"
  run_command $h "mv ~/home/vagrant/data_H/* ~/.qosd/data/"
  
done

echo "***需要手动启动当前及目标服务器qosd程序***"
exit 1

