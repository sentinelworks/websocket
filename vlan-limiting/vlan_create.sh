#!/bin/bash

#ip link add link eth0 name eth0.8 type vlan id 8

declare -a interfaces=("eth1" "eth2" "eth3" "eth4")

for k in "${interfaces[@]}"
do
    for i in {1..512}
    do
        echo "/sbin/ip link add link eth1 name $k.$i type vlan id $i"
    done
done

