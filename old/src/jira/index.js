import { createFormatter } from "../formatter/index.js";
import JiraClient from "jira-client";
import { config } from "../config/index.js";

export class Jira {
  constructor() {
    const { hostname, username, password } = config.getSettings();
    this.client = new JiraClient({
      protocol: "https",
      host: hostname,
      username,
      password,
      apiVersion: "2",
      strictSSL: true,
    });
  }

  async getIssues(ids, format) {
    // noinspection JSUnresolvedFunction
    const results = await Promise.all(
      ids.map((id) => this.client.findIssue(id))
    );
    const formatter = createFormatter(format);
    console.log(formatter.format(results));
  }
}
