#!/bin/bash
# @titulo: azure template vm linux multi nic
# @sector: DevOps
# @author myu1d157h0u54nd
# @version 1.0
# @date 03/2017
# @right azure subscription
# touch template-azure-client-env.sh && chmod 755 template-azure-client-env.sh && vim template-azure-client-env.sh
# run example
# ./template-azure-client-env.sh > audit-logs.log

read -n1 -r -p 'Press ENTER key to continue...DEPLOY...' key
echo -e '\n'

# set variable------------------------------------------------------------------
sufijoNodo="nodo02"
sufijoNET="NET-"

ResourceGroup="appExample-client-env"
location="southcentralus"
VnetName="Vnet2"
VnetName=$sufijoNET$VnetName
VnetPrefix="10.2.0.0/16"

VnetSubnetAName="VnetSubnetA"
VnetSubnetAName="default"
VnetSubnetAPrefix="10.2.0.0/24"
NSG_A="NSG_A"
NSG_A=$sufijoNodo$NSG_A
NICnameA="NIC_A"
NICnameA=$sufijoNodo$NICnameA
NICAipPrivate="10.2.0.5"
ipPublicAname="ipPublicA"
ipPublicAname=$sufijoNodo$ipPublicAname

VnetSubnetBName="VnetSubnetB"
VnetSubnetBName=$sufijoNodo$VnetSubnetBName
VnetSubnetBPrefix="10.2.5.0/24"
NSG_B="NSG_B"
NSG_B=$sufijoNodo$NSG_B
NICnameB="NIC_B"
NICnameB=$sufijoNodo$NICnameB
NICBipPrivate="10.2.5.5"
ipPublicBname="ipPublicB"
ipPublicBname=$sufijoNodo$ipPublicBname

vmName="vmName"
vmName=$sufijoNodo$vmName
diskVMname="diskVM"
diskVMname=$sufijoNodo$diskVMname
storageSKU="Standard_LRS"
#storageSKU={Premium_LRS, Standard_GRS, Standard_LRS, Standard_RAGRS, Standard_ZRS}
size="Standard_A2_v2"
imageVM="Canonical:UbuntuServer:16.04-LTS:latest"
adminUsername="sre"
adminPassword="ToDoChange"

read -n1 -r -p 'Press ENTER key to continue...create a resource group...' key
echo -e '\n'
# resource group and location --------------------------------------------------
# create a resource group
# az group create --name $ResourceGroup --location $location

read -n1 -r -p 'Press ENTER key to continue...Create vnet...' key
echo -e '\n'
# network (vnet)-------------------------------------------------------------------
# Create the virtual network
az network vnet create \
--resource-group $ResourceGroup \
--name $VnetName --address-prefix $VnetPrefix

# network (a)-------------------------------------------------------------------
read -n1 -r -p 'Press ENTER key to continue...Create subnet a...' key
echo -e '\n'
# Create subnet (a)
az network vnet subnet create \
--vnet-name $VnetName \
--resource-group $ResourceGroup \
--name $VnetSubnetAName \
--address-prefix $VnetSubnetAPrefix

read -n1 -r -p 'Press ENTER key to continue...Create nsg a...' key
echo -e '\n'
# creates a network security group (a)
az network nsg create \
--resource-group $ResourceGroup \
--name $NSG_A

# ipPublic (a)
az network public-ip create \
--name $ipPublicAname \
--resource-group $ResourceGroup \
--allocation-method Static

read -n1 -r -p 'Press ENTER key to continue...Create NICs a...' key
echo -e '\n'
# Create two NICs: nic (a) + public ip
az network nic create \
--resource-group $ResourceGroup \
--name $NICnameA \
--vnet-name $VnetName \
--subnet $VnetSubnetAName \
--network-security-group $NSG_A \
--private-ip-address $NICAipPrivate \
 --public-ip-address $ipPublicAname

# network (b)-------------------------------------------------------------------
read -n1 -r -p 'Press ENTER key to continue...Create subnet b...' key
echo -e '\n'
# Create subnet (b)
az network vnet subnet create \
--vnet-name $VnetName \
--resource-group $ResourceGroup \
--name $VnetSubnetBName \
--address-prefix $VnetSubnetBPrefix

read -n1 -r -p 'Press ENTER key to continue...Create nsg b...' key
echo -e '\n'
# creates a network security group (b)
az network nsg create \
--resource-group $ResourceGroup \
--name $NSG_B

# ipPublic (b)
# az network public-ip create \
# --name $ipPublicAname \
# --resource-group $ResourceGroup \
# --allocation-method Static

read -n1 -r -p 'Press ENTER key to continue...Create NICs b...' key
echo -e '\n'
# Create two NICs: nic (b) + public ip
az network nic create \
--resource-group $ResourceGroup \
--name $NICnameB \
--vnet-name $VnetName \
--subnet $VnetSubnetBName \
--network-security-group $NSG_B \
--private-ip-address $NICBipPrivate
# \ --public-ip-address $ipPublicA

read -n1 -r -p 'Press ENTER key to continue...Create a VM and attach the NICs NICAName NICBName...' key
echo -e '\n'
# Create a VM and attach the NICs ----------------------------------------------
az vm create \
--resource-group $ResourceGroup \
--name $vmName \
--os-disk-name $diskVMname \
--image $imageVM \
--storage-sku $storageSKU \
--size $size \
--admin-username $adminUsername \
--admin-password $adminPassword \
--authentication-type password \
--nics $NICnameA $NICnameB


