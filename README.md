# Jira Slack Standup

CLI tool, which takes Jira key(s) as input and outputs to stdout and copies to clipboard formatted output for Slack.

## Installation

```shell
git clone git+https://github.com/jmaver-plume/jira-slack-standup.git
cd jira-slack-standup
npm install -g .
```

## CLI Commands

### Config

#### Set

```shell
jss config set

# Output
prompt: username: <username> [optional]
prompt: password: <hidden-password> [optional]
prompt: hostname: <hostname> [required]
```

#### Get

```shell
jss config get

# Output
{
  username: '<username>',
  password: '<password>',
  hostname: '<hostname>'
}
```

### Issue

#### Get

```shell
jss issue get JRA-9 JRA-10

# Output
[JRASERVER-9](https://<hostname>/browse/JRASERVER-9): <jira-summary>
[JRASERVER-10](https://<hostname>/browse/JRASERVER-10): <jira-summary>
```

### Help

Every command and subcommand has a help option.

```shell
jss issue --help
jss config get --help
...
```
