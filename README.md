# Jira enrich

Jira enrich CLI tool used for my Alfred workflow.  
Converts `<jira-key> flags` into `<slack-emoticons> <jira-key>: <jira-summary>`

## Example

```
# input
JIRA-1 b
JIRA-2

# output (in clipboard)
:building: JIRA-1: Summary of JIRA-1 task
JIRA-2: Summary of JIRA-2 task
```


## Usage

1. Build binary `go build enrich.go`
2. Help `./enrich --help`
3. Run `./enrich -keys="$KEYS" -username="$USERNAME" -password="$PASSWORD" -baseUrl="$BASE_URL"`

