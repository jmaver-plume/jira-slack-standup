# Jira enrich for Alfred and Slack

CLI tool, which takes Jira key as input and returns formatted output for Slack.

## Installation

```shell
git clone git+https://github.com/jmaver-plume/jira-enrich-alfred.git
cd jira-enrich-alfred
npm install -g .
```

## CLI Commands

### Config

Path of configuration file is `~/.jira/config`.

#### Set

```shell
jira config set 

# Output
prompt: username: <username>
prompt: password: <hidden-password>
prompt: hostname: <hostname>
```

#### Get

```shell
jira config get

# Output
{
  username: '<username>',
  password: '<password>',
  hostname: '<hostname>'
}
```

### Issue

#### Get one

```shell
jira issue get JIRASUP-20 -f slack

# Output
[JIRASUP-20](https://<hostname>/browse/JIRASEUP-20): <jira-summary>
```

#### Get many

```shell
jira issue get JIRASUP-20 JIRASUP-21 -f slack

# Output
[JIRASUP-20](https://<hostname>/browse/JIRASEUP-20): <jira-summary>
[JIRASUP-21](https://<hostname>/browse/JIRASEUP-21): <jira-summary>
```
