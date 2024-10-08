#!/bin/bash
BIN=../../../zookeeper
mkdir -p zk_conf/zoo_$5
printf "tickTime=2000\ndataDir=/tmp/zookeeper/zk_$5\nclientPort=$4\ninitLimit=5\nsyncLimit=2\nserver.1=localhost:2888:$1\nserver.2=localhost:2889:$2\nserver.3=localhost:2890:$3" >> zk_conf/zoo_$5/zoo.cfg
mkdir -p /tmp/zookeeper/zk_$5
printf $5 >> /tmp/zookeeper/zk_$5/myid
${BIN}/bin/zkServer.sh --config zk_conf/zoo_$5 start-foreground
