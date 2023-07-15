import { Config } from "../adapters/out/persistence/config.js";

export function getConfigUseCase() {
  const config = new Config();
  return config.get();
}
