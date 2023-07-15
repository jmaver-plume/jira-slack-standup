import { Config } from "../adapters/out/persistence/config.js";
import { copyToClipboard } from "../adapters/out/clipboard/copy.js";

export function findJiraIssuesUseCase(ids) {
  const config = new Config();
  const { username, password, hostname } = config.get();

  const jira = new Jira(username, password, hostname);
  const issues = jira.findIssues(ids);

  const formattedIssues = issues.map((issue) => {
    const path = `https://${hostname}/browse/${issue.key}`;
    return `[${issue.key}](${path}): ${issue.fields.summary}`;
  });

  copyToClipboard(formattedIssues);
  return formattedIssues;
}
