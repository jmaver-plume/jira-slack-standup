import { mkdirp } from "mkdirp";
import fs from "fs";
import path from "path";
import os from "os";

class ConfigNotError extends Error {}

export class Filesystem {
  getConfig() {
    const filePath = this.#getFilePath();
    if (!fs.existsSync(filePath)) {
      throw new ConfigNotError();
    }

    const serializedConfig = fs.readFileSync(filePath, "utf-8");
    return JSON.parse(serializedConfig);
  }

  setConfig(config) {
    const directoryPath = this.#getDirectoryPath();
    mkdirp.sync(directoryPath);

    const filePath = this.#getFilePath();
    const serializedConfig = JSON.stringify(config, null, 2);
    return fs.writeFileSync(filePath, serializedConfig);
  }

  #getFilePath() {
    const directoryPath = this.#getDirectoryPath();
    return path.join(directoryPath, "config");
  }

  #getDirectoryPath() {
    return path.join(os.homedir(), ".jss");
  }
}
