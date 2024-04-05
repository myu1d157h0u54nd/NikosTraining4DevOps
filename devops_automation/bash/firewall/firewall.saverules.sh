#!/bin/bash
iptables-save > /root/bkp-firewall-rules-$(date +"%Y-m%m-%d")
