#!/usr/bin/env node

import commander from "commander";
import { setConfigCommand } from "./set-config-command";
import { getConfigCommand } from "./get-config-command";
import { findJiraIssuesCommand } from "./find-jira-issues-command";

const program = new commander.Command();

setConfigCommand(program);
getConfigCommand(program);
findJiraIssuesCommand(program);

program.parse(process.argv);
