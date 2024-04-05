#!/bin/bash


iptables -A FORWARD -i enp0s3 -d 10.25.59.55 -p tcp --dport 8085 -j ACCEPT
iptables -t nat -A PREROUTING -i enp0s3 -p tcp --dport 8085 -j DNAT --to 10.25.59.55:8085

iptables -A FORWARD -i enp0s3 -d 10.25.59.55 -p tcp --dport 8086 -j ACCEPT
iptables -t nat -A PREROUTING -i enp0s3 -p tcp --dport 8086 -j DNAT --to 10.25.59.55:8086

iptables -A FORWARD -i enp0s3 -d 10.25.59.55 -p tcp --dport 8081 -j ACCEPT
iptables -t nat -A PREROUTING -i enp0s3 -p tcp --dport 8081 -j DNAT --to 10.25.59.55:8081

iptables -A FORWARD -i enp0s3 -d 10.25.59.55 -p tcp --dport 1119 -j ACCEPT
iptables -t nat -A PREROUTING -i enp0s3 -p tcp --dport 1119 -j DNAT --to 10.25.59.55:1119

