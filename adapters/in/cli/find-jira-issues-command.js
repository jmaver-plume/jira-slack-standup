import { findJiraIssuesUseCase } from "../../../application/find-jira-issues-use-case.js";

export function findJiraIssuesCommand(program) {
  program
    .command("issue")
    .command("get <ids...>")
    .action(async (ids) => {
      // TODO: Verify if this is always an array
      const issues = findJiraIssuesUseCase(ids);
      console.log(issues);
    });
}
