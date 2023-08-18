#!/usr/bin/env node

import commander from "commander";
import getIssuesCommand from "./get-issues-command.js";
import setConfigCommand from "./set-config-command.js";
import getConfigCommand from "./get-config-command.js";

const program = new commander.Command();
const configCommand = program.command("config");

getIssuesCommand(program);
setConfigCommand(configCommand);
getConfigCommand(configCommand);

program.parse(process.argv);
