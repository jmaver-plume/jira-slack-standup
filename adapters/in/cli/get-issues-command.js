import { getJiraTicketDescriptions } from "../../../domain/get-jira-ticket-descriptions.js";

export default function getIssuesCommand(program) {
  program
    .command("issue")
    .command("get <ids...>")
    .action(async (ids) => {
      const issues = await getJiraTicketDescriptions(ids);
      console.log(issues);
    });
}
