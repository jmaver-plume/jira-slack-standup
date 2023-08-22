import { getConfig } from "./get-config.js";
import { Jira } from "../adapters/out/jira/index.js";
import { writeToClipboard } from "../adapters/out/clipboard/index.js";

function formatIssue(issue, hostname) {
  const path = `https://${hostname}/browse/${issue.key}`;
  return `[${issue.key}](${path}): ${issue.fields.summary}`;
}

function formatIssues(issues, hostname) {
  return issues.map((issue) => formatIssue(issue, hostname)).join("\n");
}

export async function getJiraTicketDescriptions(ids) {
  const { username, password, hostname } = getConfig();
  const jira = new Jira(hostname, {username, password});
  const issues = await jira.getIssues(ids);
  const formattedIssues = formatIssues(issues, hostname);
  const serializedFormattedIssues = formattedIssues.toString();
  writeToClipboard(serializedFormattedIssues);
  return formattedIssues;
}
