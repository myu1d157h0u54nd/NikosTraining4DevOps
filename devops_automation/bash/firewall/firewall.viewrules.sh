#!/bin/bash

iptables -L -n -v

iptables -t nat -L

