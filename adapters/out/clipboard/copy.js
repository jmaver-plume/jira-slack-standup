import clipboard from "clipboardy";

export function copyToClipboard(data) {
  clipboard.writeSync(data);
}
