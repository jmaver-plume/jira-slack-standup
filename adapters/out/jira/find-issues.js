import JiraClient from "jira-client";

export class Jira {
  constructor(username, password, hostname) {
    this.client = new JiraClient({
      protocol: "https",
      host: hostname,
      username,
      password,
      apiVersion: "2",
      strictSSL: true,
    });
  }

  async findIssues(ids) {
    const promises = ids.map((id) => this.client.findIssue(id));
    return await Promise.all(promises);
  }
}
