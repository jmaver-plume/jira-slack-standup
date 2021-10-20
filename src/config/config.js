import fs from "fs";
import mkdirp from "mkdirp";
import os from "os";
import path from "path";

export class Config {
  directory = ".jira";
  filename = "config";

  getDirectoryPath() {
    return path.join(os.homedir(), this.directory);
  }

  getFilePath() {
    const directoryPath = this.getDirectoryPath();
    return path.join(directoryPath, this.filename);
  }

  getSettings() {
    return JSON.parse(fs.readFileSync(this.getFilePath(), "utf-8"));
  }

  setSettings(settings) {
    const directoryPath = this.getDirectoryPath();
    mkdirp.sync(directoryPath);
    return fs.writeFileSync(
      this.getFilePath(),
      JSON.stringify(settings, null, 2)
    );
  }
}
