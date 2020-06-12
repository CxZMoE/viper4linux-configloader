# viper4linux-configloader
A config loader for Viper4Linux - An Adaptive Digital Sound Processor

# Description

```
===========================================
|         viper4linuxconfigloader         |
|                                         |
|                 welcome                 |
|                                         |
| Description:                            |
|                                         |
| this program is a loader program of the |
| project: viper4linux on github,for the  |
| convinience of switching config files.  |
|                                         |
| This program does not contain any codes |
| from the viper4 team or any other teams,|
| and will be pubulish for free with MIT  |
| licence.                                |
|                                         |
| My github: https://github.com/CxZMoE    |
| Bug reports are welcomed.               |
|                                         |
===========================================
```

# Installation

## Debian Package

the program installed by debian package have config files provided by viper4linux author and community which is located in:
`/usr/local/viper4linux-configloader/configs`,you can use this path to fill in the settings section in the Installation Guide below.

```shell
wget https://github.com/CxZMoE/viper4linux-configloader/releases/download/v1.0/viper4linux-configloader-linux-amd64.deb
sudo dpkg -i viper4linux-configloader-linux-amd64.deb
```

## Manual

1. download [released](https://github.com/CxZMoE/viper4linux-configloader/releases/tag/v1.0) executable file.
2. change the filename to whatever you want it to be.
```shell
mv viper4linux-configloader-linux-amd64 viper4linux-configloader
```
3. copy it to your bin directory(like ~/.local/bin|~/.bin|/usr/bin,etc),make sure the directory is included in your environment path.
``` shell
# Check environment path.
echo $PATH
# /home/xxx/go/bin:/home/xxx/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin
# make sure the your bin directory is included in your environment path

sudo cp viper4linux-configloader /usr/bin
```
3. run 
```shell
viper4linux-configloader
```

# Usage

run program
```shell
viper4linux-configloader

```

1. You need to configure settings first when first run.
- this step you need to set where you put config files,like: `/home/xxx/myconfigs/`
  - You can use `/usr/local/viper4linux-configloader/configs` if you installed from debian package.

- then set where the audio.conf is,for my case,it's  `/home/cxzmoe/.config/viper4linux/audio.conf`
  - Notmally in `~/.config/viper4linux` if you follow the installation guide of viper4linux.


`tip: if you have an empty config dicrectory, you need to fill it first before run this program.`

when you configured settings, and config directory is not empty,the program should look like this:
```shell
Viper4Linux found at: /home/cxzmoe/.local/bin/viper


===========================================
|         viper4linuxconfigloader         |
|                                         |
|                 welcome                 |
|                                         |
| Description:                            |
|                                         |
| this program is a loader program of the |
| project: viper4linux on github,for the  |
| convinience of switching config files.  |
|                                         |
| This program does not contain any codes |
| from the viper4 team or any other teams,|
| and will be pubulish for free with MIT  |
| licence.                                |
|                                         |
| My github: https://github.com/CxZMoE    |
| Bug reports are welcomed.               |
|                                         |
===========================================

Quit with 'Ctrl+C'.

0.[000-default-audio.conf]
1.[loque-akg_k67_bass-audio.conf]
2.[matis-FDS_soft-audio.conf] [With IRS]
3.[matis-FDS_v1-audio.conf] [With IRS]
4.[thepbone-clear_bass-audio.conf] [With IRS]
5.[topjor-srs_2_1-audio.conf] [With IRS]
6.[unequaled86-cCc-audio.conf] [With IRS]

Load config of number? 

```

2. Select the config file you want to switch to and press enter.

3. ENJOY!

# QA
1. How to clear my config?
simply `rm -rf ~/.config/viper4linux-configloader`

2. Where is the preprovided config files?
`/usr/local/viper4linux-configloader/configs`

3. Where is the audio.conf?
Notmally in `~/.config/viper4linux` if you follow the installation guide of viper4linux.
