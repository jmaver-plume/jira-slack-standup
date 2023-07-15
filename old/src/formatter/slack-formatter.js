import { config } from "../config/index.js";

export class SlackFormatter {
  format(input) {
    if (Array.isArray(input)) {
      return this.formatList(input);
    }

    return this.formatItem(input);
  }

  formatList(items) {
    return items.map((item) => this.formatItem(item)).join("\n");
  }

  formatItem(item) {
    const host = config.getSettings().hostname;
    const path = `https://${host}/browse/${item.key}`;
    return `[${item.key}](${path}): ${item.fields.summary}`;
  }
}
