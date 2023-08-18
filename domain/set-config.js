import { Filesystem } from "../adapters/out/filesystem/index.js";

export function setConfig(config) {
  const filesystem = new Filesystem();
  filesystem.setConfig(config);
}
