# ec2ssh

## ec2sshとは
踏み台サーバ(bastion)から、簡単に各EC2サーバーへログインするためのツールです。

## 前提条件
以下が使用できる必要があります。
OSは、AmazonLinux2,CentOS7で動作確認済みです。

- awscli
- peco
- bash

## 使い方
```
Usage: ec2ssh [SSHUser]
   or: ec2ssh [SSHUser] [GrepWord]
```
※そのうちDEMO動画を載せます。

## Install
念のため`awscli`と`peco`についてもインストール手順を記載します。
AmazonLinux2であれば、`awscli`はデフォルトでインストールされていると思います。

#### Install awscli:
```bash
$ sudo yum install epel-release
$ sudo yum install python-pip
$ sudo pip install pip --upgrade
$ pip install awscli --user
```

####  Install peco:
最新バージョンを確認します。
https://github.com/peco/peco/releases
-> v0.5.3(as of October 26, 2018)
```bash
$ cd ~/
$ mkdir -p local/src
$ cd local/src
$ wget https://github.com/peco/peco/releases/download/v0.5.3/peco_linux_amd64.tar.gz
$ tar -C ~/local -zxvf peco_linux_amd64.tar.gz

# Move decompressed peco command binary to /usr/local/bin/
$ sudo mv ~/local/src/peco_linux_amd64/peco /usr/local/bin/

# Check
$ which peco
```

### Install ec2ssh
```bash
$ cd ~/
$ mkdir -p local/src
$ cd local/src
$ wget https://raw.githubusercontent.com/tosasaki/ec2ssh/master/ec2ssh
# Move ec2ssh command binary to /usr/local/bin/
$ sudo mv ~/local/src/ec2ssh /usr/local/bin/

# Check
$ which ec2ssh
```
