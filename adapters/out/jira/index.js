import JiraClient from "jira-client";

export class Jira {
  #client;
  constructor(hostname, options = {}) {
    this.#client = new JiraClient({
      protocol: "https",
      host: hostname,
      username: options.username,
      password: options.password,
      apiVersion: "2",
      strictSSL: true,
    });
  }

  async getIssues(ids) {
    // noinspection JSUnresolvedFunction
    return await Promise.all(ids.map((id) => this.#client.findIssue(id)));
  }
}
