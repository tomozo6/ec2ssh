# Overview
ec2ssh is a tool that can easily ssh login to AWS EC2.

# Install
## Homebrew (macOS and Linux)
```bash
$ brew install tomozo6/tap/ec2ssh
```
# Requirement
## awsci
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
[peco Installation](https://github.com/peco/peco#installation)

# Usage
```
Usage:
  ec2ssh [-g grepword] [-s] [-u user] ...

Description:
  ec2ssh is a tool that can easily ssh login to AWS EC2.

Options:
  -g specify the word you want to grep.
  -s use SSM SessionManager. (use the InstanceID instead of IpAddress.)
  -u specify the user you want to ssh. (default: ec2-user)
  -h show help.
```

## SSM SessionManager
To use SSM SessionManager, make the following settings.

[Enable SSH connections through Session Manager](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/session-manager-getting-started-enable-ssh-connections.html)

# Licence
MIT

# Author
tomozo6