#!/bin/bash

#ip link add link eth0 name eth0.8 type vlan id 8
# vconfig add eth1 10
# vconfig rem eth1.10

# ip a add 192.168.10.2/24 dev eth1.10
# ifconfig eth1.10 up

declare -a interfaces=("eth1" "eth2" "eth3" "eth4")

for k in "${interfaces[@]}"
do
    j=${k:3:1}
    for i in {1..200}
    do
        echo "/sbin/ip a add 192.168.$i.$j/24 $k.$i"
        /sbin/ip a add 192.168.$i.$j/24 $k.$i
        ifconfig $k.$i up
    done
done

