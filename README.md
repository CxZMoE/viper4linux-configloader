# viper4linux-configloader
A config loader for Viper4Linux - An Adaptive Digital Sound Processor

# Description

```
===========================================
|             v4dconfigloader             |
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

1. download [released](https://github.com/CxZMoE/viper4linux-configloader/releases/tag/v1.0) executable file
2. change the filename to whatever you want it to be.
```shell
mv viper4linux-configloader-linux-amd64 viper-loader
```
3. copy it to your bin directory(like ~/.local/bin|~/.bin|/usr/bin,etc),make sure the directory is included in your environment path.
``` shell
// Check environment path.
echo $PATH
/*
...
/usr/bin
*/
sudo cp viper-loader /usr/bin
```
3. run 
```shell
viper-loader
```

# Usage
1. You need to configure settings first when first run.
- this step you need to set where you put config files,like: `/home/xxx/myconfigs/`
- then set where the audio.conf is,for my case,it's  `/home/cxzmoe/.config/viper4linux/audio.conf`

`tip: if you have an empty config dicrectory, you need to fill it first before run this program.`

when you configured settings, and config directory is not empty,the program should look like this:
```shell
Viper4Linux found at: /home/cxzmoe/.local/bin/viper


===========================================
|             v4dconfigloader             |
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

Please enter number of -1 to quit, quit with 'Ctrl+C' will cause viper stop.

0.[000-default-audio.conf]
1.[loque-akg_k67_bass-audio.conf]
2.[matis-FDS_soft-audio.conf] [With IRS]
3.[matis-FDS_v1-audio.conf] [With IRS]
4.[thepbone-clear_bass-audio.conf] [With IRS]
5.[topjor-srs_2_1-audio.conf] [With IRS]
6.[unequaled86-cCc-audio.conf] [With IRS]

Load config of number? ( -1 to quit ) 

```

2. Select the config file you want to switch to and press enter.

3. Input -1 and then Enter to quit, if you use CTRL+C to quit, the viper will be killed by SIGKILL.

4. ENJOY!
