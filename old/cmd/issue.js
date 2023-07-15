#!/usr/bin/env node

import commander from "commander";
import { Format } from "../src/formatter/index.js";
import prompt from "prompt";
import { config } from "../src/config/index.js";
import { Jira } from "../src/jira/index.js";

const FormatOption = new commander.Option("-f, --format <format>", "").choices([
  Format.Alfred,
  Format.Slack,
]);

const program = new commander.Command();

const configCmd = program.command("config");

configCmd.command("set").action(() => {
  prompt.start();
  return new Promise((resolve, reject) =>
    prompt.get(
      {
        properties: {
          username: {
            required: true,
          },
          password: {
            hidden: true,
            required: true,
          },
          hostname: {
            required: true,
          },
        },
      },
      (err, settings) => {
        if (err) {
          return reject(err);
        }
        config.setSettings(settings);
      }
    )
  );
});

configCmd.command("get").action(() => console.log(config.getSettings()));

const issue = program.command("issue");

issue
  .command("get <ids...>")
  .addOption(FormatOption)
  .action(async (ids, options) => {
    const jira = new Jira();
    await jira.getIssues(ids, options.format);
  });

program.parse(process.argv);
