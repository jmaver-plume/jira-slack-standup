import { Filesystem } from "../adapters/out/filesystem/index.js";

export function getConfig() {
  const filesystem = new Filesystem();
  return filesystem.getConfig();
}
