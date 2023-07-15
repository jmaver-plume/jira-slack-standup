export class AlfredFormatter {
  format(arg) {
    if (Array.isArray(arg)) {
      return JSON.stringify(this.formatList(arg), null);
    }

    return JSON.stringify(this.formatItem(arg), null);
  }

  formatList(items) {
    return {
      items: items.map((item) => this.formatItem(item)),
    };
  }

  formatItem(item) {
    return {
      uid: item.fields.summary,
      type: "",
      title: item.fields.summary,
      subtitle: `${item.fields.status.name}: [${item.fields.assignee.displayName}]`,
      arg: item.fields.summary,
      autocomplete: item.fields.summary,
      icon: {
        type: item.fields.issuetype.iconUrl,
        path: "",
      },
    };
  }
}
