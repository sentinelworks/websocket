#!/bin/bash

#/sbin/ip link add link eth0 name eth0.8 type vlan id 8
#/sbin/ip -d link show eth1.48

interface="eth1"

for i in {1..5}
do
    echo "/sbin/ip -d link show ${interface}.$i"
done

