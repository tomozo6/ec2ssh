 #### Read this in other languages

<kbd>[<img title="日本語" alt="日本語" src="https://cdn.staticaly.com/gh/hjnilsson/country-flags/master/svg/jp.svg" width="22">](/README.ja.md)</kbd>

# ec2ssh

`ec2ssh` is a tool that can easily ssh login to AWS EC2.

Finally, it's a SSH wrapper tool that just generates and executes the following ssh command.

`ssh ${user}@${LocalIpAddress}` or `ssh ${user}@${InstanceID}`

Therefore, the SSH configuration file is also applied.

(The SSH config file is usually located in `~/.ssh/config`.)

## Install

### MacOS

```bash
brew install tomozo6/tap/ec2ssh
```

### Linux

```bash
wget https://github.com/tomozo6/ec2ssh/releases/download/v0.0.3/ec2ssh_Linux_x86_64.tar.gz
tar zxvf ec2ssh_Linux_x86_64.tar.gz
chmod +x ./ec2ssh
sudo mv ./ec2ssh /usr/local/bin/ec2ssh
```

## Requirement

### awscli

Required if you want to make an SSH connection through Session Manager.

Please refer to the following URL for the installation method.
[Installing, updating, and uninstalling the AWS CLI](https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-chap-install.html)

### Session Manager plugin

Required if you want to make an SSH connection through Session Manager.

Also, if you want to make an SSH connection through Session Manager, you need the Session Manager plug-in version `1.1.23.0` or higher.

If the old plugin is installed or the plugin is not installed in the first place, please install the latest version.

Please refer to the following URL for the installation method.

[(Optional) Install the Session Manager plugin for the AWS CLI](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html)

## Usage

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

### Normal SSH connection

[Example]

```bash
ec2ssh
```

Since `ec2ssh` uses the local IP address of EC2 as the host name of the SSH connection destination, it is assumed to be used on the bastion server.

To perform multi-stage SSH from your PC via the bastion server, you need to customize the SSH configuration file on your PC. (Not explained here)

### SSH connection through Session Manager

[Example]

```bash
ec2ssh -s
```

If you add the option `-s` to`ec2ssh`, the instance ID of EC2 will be the host name of the SSH connection destination, so you can connect with Session Manager.

You also need to add the following to your PC's SSH config file.

```bash
# SSH over Session Manager
host i-* mi-*
    ProxyCommand sh -c "aws ssm start-session --target %h --document-name AWS-StartSSHSession --parameters 'portNumber=%p'"
```

For details, refer to "To enable SSH connections through Session Manager" and "2. On the local machine from which you want to connect to a managed instance using SSH, do the following:" of the following URL.

[Step 8: (Optional) Enable SSH connections through Session Manager](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/session-manager-getting-started-enable-ssh-connections.html)

### Use ec2ssh config file

ec2ssh can describe options in the configuration file. By default, `~/.ec2ssh.yaml` is automatically read as a configuration file.
You can also use the option `--config` to load any configuration file.

[example]

```yaml
session-manager: true
ssh-user: tomozo6
```

## Licence

MIT

## Author

tomozo6
