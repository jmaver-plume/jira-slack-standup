import prompt from "prompt";
import { config } from "../src/config/index.js";

export function setConfigCommand(program) {
  program
    .command("config")
    .command("set")
    .action(async () => {
      prompt.start();
      const { username, password, hostname } = await new Promise(
        (resolve, reject) =>
          prompt.get({
            properties: {
              username: {
                required: true,
              },
              password: {
                hidden: true,
                required: true,
              },
              hostname: {
                required: true,
              },
            },
          })
      );

      setConfigUseCase(username, password, hostname);
    });
}
