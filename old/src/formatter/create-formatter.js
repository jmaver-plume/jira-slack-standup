import { SlackFormatter } from "./slack-formatter.js";
import { AlfredFormatter } from "./alfred-formatter.js";
import { DefaultFormatter } from "./default-formatter.js";
import { Format } from "./enums.js";

export function createFormatter(format) {
  switch (format) {
    case Format.Slack:
      return new SlackFormatter();
    case Format.Alfred:
      return new AlfredFormatter();
    default:
      return new DefaultFormatter();
  }
}
