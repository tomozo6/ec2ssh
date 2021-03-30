# 概要
`ec2ssh` は、AWS EC2へのSSHログインを簡単にするためのツールです。

# 詳細
最終的に以下のようなsshコマンドを生成して実行しているだけです。

`ssh ${user}@${LocalIpAddress}` or `ssh ${user}@${InstanceID}`


# インストール方法
## Homebrew (macOS and Linux)
```bash
$ brew install tomozo6/tap/ec2ssh
```
# 前提条件
## awsci
インストール方法は以下のURLを参考にしてください。

[Installing, updating, and uninstalling the AWS CLI](https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-chap-install.html)

## peco
### macOS (Homebrew)
```bash
brew install peco
```
### Debian and Ubuntu based distributions (APT)
```bash
apt install peco
```
`peco`の詳細は以下のURLを参考にしてください。

[peco Installation](https://github.com/peco/peco#installation)

## Session Manager plugin
Session Manager を使用する場合は必要です。

また Session Manager を通してSSH接続をする場合は、バージョン `1.1.23.0` 以上の Session Managerプラグインが必要です。

古いプラグインがインストールされいたり、そもそもプラグインがインストールされていない場合は、最新版をインストールして下さい。

インストール方法は以下のURLを参考にしてください。

[(オプション) AWS CLI 用の Session Manager plugin をインストールする](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html)


# 使い方
```
使い方:
  ec2ssh [-g grepword] [-s] [-u user] ...

説明:
  ec2ssh is a tool that can easily ssh login to AWS EC2.

オプション:
  -g Grepしたいワードを指定してください。
  -s SSMセッションマネージャーを使用してログインしたい場合に指定してください。
     (IPアドレスではなくインスタンスIDでSSHしようと試みます)
  -u ログインしたいSSHユーザーを指定します. (default: ec2-user)
  -h ヘルプを表示します。
```
## 通常のSSH接続をする
`ec2ssh`はEC2のローカルIPアドレスをSSH接続先のホスト名とするため、踏み台サーバーでの使用を想定しています。

自分のPCから、踏み台サーバーを経由して多段SSHをするには、自分のPCのSSH設定ファイルをカスタマイズする必要があります。(ここでは説明しません)


## Session Managerを通してSSH接続をする
自分のPCのSSH設定ファイルに以下を追記します。

(SSH設定ファイルは通常`~/.ssh/config`にあります。)
```bash
# SSH over Session Manager
host i-* mi-*
    ProxyCommand sh -c "aws ssm start-session --target %h --document-name AWS-StartSSHSession --parameters 'portNumber=%p'"
```
詳細は以下のURLの

「Session Managerを通して SSH 接続を有効にする」の 「2. SSH を使用してマネージドインスタンスに接続するローカルマシンで、次の手順を実行します。」
を参考にしてください。

[ステップ 8: (オプション) Session Manager を通して SSH 接続を有効にする](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/session-manager-getting-started-enable-ssh-connections.html)

# ライセンス
MIT

# 著者
tomozo6