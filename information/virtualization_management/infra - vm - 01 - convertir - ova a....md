
# Metodos para convertir un OVa en otros formatos para cloud

- incluye OVA 2 VHDX for Hyper-V
- incluye OVA 2 VHDX for Hyper-V
- incluye OVA 2 VHDX for Hyper-V
- incluye OVA 2 VHDX for Hyper-V


## OVA 2 VHDX for Hyper-V

### Required software
 - Microsoft Virtual Machine Converter 3.0 available here.
 - Your OVA export unzipped into a folder.
 - windows pc

### Preparations
 - Download and install Microsoft Virtual Machine Converter 3.0.
 - Download and install Windows Management Framework 3.0

### Microsoft Virtual Machine Converter 3.0
https://www.microsoft.com/en-us/download/details.aspx?id=42497

### Windows Management Framework 3.0
https://www.microsoft.com/en-us/download/details.aspx?id=34595
### Convert The Image

Next we will need to convert our VMware Image in order for Hyper-V to run it. This can be done using Powershell as seen below.

# Import the Converter Powershell Module
Import-Module "C:\Program Files\Microsoft Virtual Machine Converter\MvmcCmdlet.psd1"

# Convert the VMware .vmdk to a Hyper-V .vhdx

# Descomprimir y extraer todos los archivos
    tar -xvf "VUAppSRV_01 - vs 3.0.1 release 6.0.4 - ova vs1.0.ova"



ConvertTo-MvmcVirtualHardDisk -SourceLiteralPath "VUAppSRV_01-vs3.vmdk" -DestinationLiteralPath "VUAppSRV_01-vs3.vhdx" -VhdType DynamicHardDisk -VhdFormat Vhdx

### Using The Image

In order to use our newly created image we must first create a new Hyper-V virtual machine or edit one that has been previously created. Important notes below.

When creating your new virtual machine, you must ensure you select “Generation 1” when choosing the generation of the virtual machine.
The converted image can be moved to any location you choose.
Once you have completed all the steps above, you should have a fully working copy of your virtual machine running on Hyper-V.

