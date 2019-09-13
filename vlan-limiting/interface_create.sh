#!/bin/bash

#ip link add link eth0 name eth0.8 type vlan id 8

interface="eth1"

for i in {1..5}
do
    #`/sbin/ip link add link eth1 name ${interface}.$i type vlan id $i`
    echo "/sbin/ip link add link eth1 name ${interface}.$i type vlan id $i"
done

