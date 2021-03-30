# Overview
ec2ssh is a tool that can easily ssh login to AWS EC2.

# Install
## Homebrew (macOS and Linux)
```
$ brew install tomozo6/tap/ec2ssh
```
# Requirement
- awscli
- peco

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

# Licence
MIT

# Author
tomozo6