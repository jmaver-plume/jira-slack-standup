import fs from "fs";
import mkdirp from "mkdirp";
import os from "os";
import path from "path";

export class Config {
  #directoryName;
  #fileName;

  constructor(directoryName = ".jira", fileName = "config") {
    this.#directoryName = directoryName;
    this.#fileName = fileName;
  }

  get() {
    const filePath = this.getFilePath();
    const file = fs.readFileSync(filePath, "utf-8");
    return JSON.parse(file);
  }

  save(username, password, hostname) {
    const directoryPath = this.getDirectoryPath();
    mkdirp.sync(directoryPath);

    const filePath = this.getFilePath();
    const stringified = JSON.stringify(
      { username, password, hostname },
      null,
      2
    );
    return fs.writeFileSync(filePath, stringified);
  }

  #getDirectoryPath() {
    const homeDir = os.homedir();
    return path.join(homeDir, this.#directoryName);
  }

  #getFilePath() {
    const directoryPath = this.getDirectoryPath();
    return path.join(directoryPath, this.#fileName);
  }
}
