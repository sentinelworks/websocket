#!/bin/bash

#/sbin/ip route add table 4294967295 to 172.28.0.0/16 dev eth3

interface="eth1"

for i in {10..50}
do
    echo "/sbin/ip route add table $i to 172.28.0.0/16 dev $interface"
done

