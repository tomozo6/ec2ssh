 #### Read this in other languages

<kbd>[<img title="日本語" alt="日本語" src="https://cdn.staticaly.com/gh/hjnilsson/country-flags/master/svg/jp.svg" width="22">](/README.ja.md)</kbd>

# ec2ssh

`ec2ssh` is a tool that can easily ssh login to AWS EC2.

Finally, it's a wrapper tool that just generates and executes the following ssh command

`ssh ${user}@${LocalIpAddress}` or `ssh ${user}@${InstanceID}`

## Install

### Homebrew (macOS and Linux)

```bash
brew install tomozo6/tap/ec2ssh
```

## Requirement

### awscli

Please refer to the following URL for the installation method.
[Installing, updating, and uninstalling the AWS CLI](https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-chap-install.html)

### peco

#### macOS (Homebrew)

```bash
brew install peco
```

#### Debian and Ubuntu based distributions (APT)

```bash
apt install peco
```

Please refer to the following URL for details of `peco`.

[peco Installation](https://github.com/peco/peco#installation)

### Session Manager plugin

You will need it if you want to use Session Manager.

Also, if you want to make an SSH connection through Session Manager, you need the Session Manager plug-in version `1.1.23.0` or higher.

If the old plugin is installed or the plugin is not installed in the first place, please install the latest version.

Please refer to the following URL for the installation method.

[(Optional) Install the Session Manager plugin for the AWS CLI](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html)

## Usage

```bash
Usage:
  ec2ssh [-g grepword] [-s] [-u user] ...

Description:
  ec2ssh is a tool that can easily ssh login to AWS EC2.

Options:
  -g specify the word you want to grep.
  -s use SSM SessionManager. (use the InstanceID instead of IpAddress.)
  -m access managed instances other than AWS. (Hybrid environment)
  -u specify the user you want to ssh. (default: ec2-user)
  -h show help.
```

### Normal SSH connection

Since `ec2ssh` uses the local IP address of EC2 as the host name of the SSH connection destination, it is assumed to be used on the bastion server.

To perform multi-stage SSH from your PC via the bastion server, you need to customize the SSH configuration file on your PC. (Not explained here)

### SSH connection through Session Manager

If you add the option `-s` to`ec2ssh`, the EC2 instance ID will be the host name of the SSH connection destination, so you can connect with Session Manager.

You also need to add the following to your PC's SSH config file.
(The SSH config file is usually located in `~ / .ssh / config`.)

```bash
# SSH over Session Manager
host i-* mi-*
    ProxyCommand sh -c "aws ssm start-session --target %h --document-name AWS-StartSSHSession --parameters 'portNumber=%p'"
```

For details, refer to "To enable SSH connections through Session Manager" and "2. On the local machine from which you want to connect to a managed instance using SSH, do the following:" of the following URL.

[Step 8: (Optional) Enable SSH connections through Session Manager](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/session-manager-getting-started-enable-ssh-connections.html)

#### Connect to a managed instance outside of the AWS environment

If you add the option `-m` to`ec2ssh`, you can access the managed instance added by hybrid activation.
Specifically, use the `ssm describe-instance-information` command to get instance information.

If the option `-m` is given, the connection through the session manager is required, so the EC2 instance ID will be the host name of the SSH connection destination regardless of the presence or absence of the option`-s`.

## Licence

MIT

## Author

tomozo6
