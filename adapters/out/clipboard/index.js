import clipboard from "clipboardy";
export function writeToClipboard(data) {
  clipboard.writeSync(data);
}
