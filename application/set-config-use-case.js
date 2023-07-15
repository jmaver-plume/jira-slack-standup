import { Config } from "../adapters/out/persistence/config.js";

export function setConfigUseCase(username, password, hostname) {
  const config = new Config();
  config.save(username, password, hostname);
}
