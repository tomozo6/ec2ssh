# ec2ssh

`ec2ssh` はAWS EC2へのSSHログインを簡単にするためのツールです。

最終的に、以下のようなsshコマンドを生成して実行しているだけのSSHラッパーツールです。

`ssh ${user}@${LocalIpAddress}` or `ssh ${user}@${InstanceID}`

そのためSSH設定ファイルも適用されます。

(SSH設定ファイルは通常`~/.ssh/config`にあります。)

## インストール方法

### Homebrew (macOS and Linux)

```bash
brew install tomozo6/tap/ec2ssh
```

## 前提条件

### awsci

Session Manager を通してSSH接続をする場合は必要です。

インストール方法は以下のURLを参考にしてください。

[Installing, updating, and uninstalling the AWS CLI](https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-chap-install.html)

### Session Manager plugin

Session Manager を通してSSH接続をする場合は必要です。
バージョン `1.1.23.0` 以上の Session Managerプラグインが必要です。

古いプラグインがインストールされいたり、そもそもプラグインがインストールされていない場合は、最新版をインストールして下さい。

インストール方法は以下のURLを参考にしてください。

[(オプション) AWS CLI 用の Session Manager plugin をインストールする](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html)

## 使い方

```bash
ec2ssh is a tool that can easily ssh login to AWS EC2.

Usage:
  ec2ssh [flags]

Flags:
      --config string     config file (default is $HOME/.ec2ssh.yaml)
  -h, --help              help for ec2ssh
  -s, --session-manager   use SSM SessionManager. (use the InstanceID instead of IpAddress.)
  -u, --ssh-user string   ssh user
  -v, --version           version for ec2ssh
```

### 通常のSSH接続をする

[実行例]

```bash
ec2ssh
```

`ec2ssh`はEC2のローカルIPアドレスをSSH接続先のホスト名とするため、踏み台サーバーでの使用を想定しています。

自分のPCから踏み台サーバーを経由して多段SSHをするには、自分のPCのSSH設定ファイルをカスタマイズする必要があります。(ここでは説明しません)

### Session Managerを通してSSH接続をする

[実行例]

```bash
ec2ssh -s
```

`ec2ssh`にオプション`-s`を付与すると、EC2のインスタンスIDをSSH接続先のホスト名とするためSession Managerでの接続が可能になります。

また自分のPCのSSH設定ファイルに以下を追記する必要があります。

```bash
# SSH over Session Manager
host i-* mi-*
    ProxyCommand sh -c "aws ssm start-session --target %h --document-name AWS-StartSSHSession --parameters 'portNumber=%p'"
```

詳細は以下のURLの

「Session Managerを通して SSH 接続を有効にする」の 「2. SSH を使用してマネージドインスタンスに接続するローカルマシンで、次の手順を実行します。」
を参考にしてください。

[ステップ 8: (オプション) Session Manager を通して SSH 接続を有効にする](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/session-manager-getting-started-enable-ssh-connections.html)

### ec2ssh設定ファイルを使用する

ec2sshは、オプションを設定ファイルに記載することが可能です。デフォルトでは`~/.ec2ssh.yaml`を設定ファイルとして自動で読み込みます。
オプション `--config`を使用して、任意の設定ファイルを読み込みことも可能です。

[設定ファイルの例]

```yaml
session-manager: true
user: tomozo6
```

## ライセンス

MIT

## 著者

tomozo6
