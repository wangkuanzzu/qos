#!/bin/bash
#脚本执行请在要导出数据的节点服务器上。参数：其他节点服务器ip

if [ ! -n "$1" ]; then
    	echo "need ips for deliver data ;example '192.168.1.1,192.168.1.2'"
    	exit 1
fi
if [ ! -n "$2" ]; then
    	echo "need type for export genesis ;type in (A,B,C,D)"
	echo "A:----从零导出到停止qosd进程时高度状态数据------"
	echo "B:----从零导出到指定高度状态数据---------"
	echo "C:----导出停止qosd进程时高度状态数据---------"
	echo "D:----导出指定高度状态数据---------"
    exit 1
fi
if [ "B" == "$2" -o "D" == "$2" ]; then
	if [ ! -n "$3" ]; then
        	echo "please set height"
		exit 1
	fi
fi

IFS="," HOSTS=($1)
for h in ${HOSTS[@]};do
  echo "***目标节点IP***：$h"
done

ip=$(ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6 | awk '{print $2}')
echo "***本机IP***$ip"

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
#0. 确保192.168.1.201启动命令：qosd start --pruning nothing 用于记录每个高度的状态保存下来。
#1. 停止192.168.1.201的qosd进程，记录当时的高度为H。导出当前H高度的genesis-H-xxxxx.json文件，替换genesis.json文件。
echo "-----操作当前服务器-----"
pkill qosd
myFile="~/.qosd/*.json"
if [ ! -f "$myFile" ];then
	rm -rf ~/.qosd/*.json
fi

echo "-----导出命令-----$2"
case $2 in 
	"A")
		echo "----从零导出到停止qosd进程高度状态数据:qosd export --for-zero-height------"
  		qosd export --for-zero-height
	;;
	"B")
		echo "----从零导出到$3高度状态数据:qosd export --for-zero-height --height $3---------"
  		qosd export --for-zero-height --height $3
	;;
	"C")
		echo "----导出到停止qosd进程高度状态数据:qosd export---------"
		qosd export 
	;;
	"D")
		echo "----导出到$3高度状态数据:qosd export --height $3---------"
		qosd export --height $3
	;;
esac

#2. 目的服务器循环
for h in ${HOSTS[@]};do
  echo $h

  #3. 将genesis-H-xxxxx.json文件传输给：192.168.1.200，同样替换原来genesis.json文件
  sshpass -pvagrant scp ~/.qosd/*.json vagrant@$h:/home/vagrant/.qosd/
  #4. 执行qosd unsafe-reset-all
  echo "-----清空$h节点数据--------------"
	
  if [ "192.168.1.200" == $h ];then
	echo "-----在192.168.1.200上执行交易，账户acc1和acc2发生变化------"
	run_command $h "$QOSCLI query account acc1"
	run_command $h "$QOSCLI query account acc2"
	
	echo "-----请在服务器$h手动执行以下命令时间30s----------"
	echo "echo '11111111'|$QOSCLI tx transfer --senders acc1,100000000000qos --receivers acc2,100000000000qos"
	#run_command $h "echo '11111111'|$QOSCLI tx transfer --senders acc1,100000000000qos --receivers acc2,100000000000qos"
	
	sleep 30s
	run_command $h "$QOSCLI query account acc1"
        run_command $h "$QOSCLI query account acc2"
  else
	sleep 30s	
  fi
  run_command $h "pkill qosd"
  run_command $h "$QOSD unsafe-reset-all"

  #5. 替换genesis文件
  run_command $h "mv ~/.qosd/*.json ~/.qosd/config/genesis.json"

done

echo "-----清空本地节点$ip数据----------"
qosd unsafe-reset-all
mv ~/.qosd/*.json ~/.qosd/config/genesis.json

#5. 手动执行qosd start，验证
echo "-----请在每个节点手动执行qosd start开启qosd进程------------------"

