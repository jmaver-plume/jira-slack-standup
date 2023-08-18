import { getConfig } from "../../../domain/get-config.js";

export default function getConfigCommand(command) {
  command.command("get").action(() => {
    const config = getConfig();
    const serializedConfig = JSON.stringify(config, null, 2);
    console.log(serializedConfig);
  });
}
