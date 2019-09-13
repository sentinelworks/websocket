#!/bin/bash

#/sbin/ip route del table 23 to 0.0.0.0/0 dev tg3

interface="eth1"

for i in {10..50}
do
    echo "/sbin/ip route add table $i to 172.28.$i.0/24 dev $interface"
done

