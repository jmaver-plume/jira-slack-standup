import prompt from "prompt";
import { setConfig } from "../../../domain/set-config.js";

export default function setConfigCommand(command) {
  command.command("set").action(() => {
    prompt.start();
    return new Promise((resolve, reject) =>
      prompt.get(
        {
          properties: {
            username: {
              required: false,
            },
            password: {
              hidden: true,
              required: false,
            },
            hostname: {
              required: true,
            },
          },
        },
        (err, config) => {
          if (err) {
            return reject(err);
          }

          const cleanConfig = {
            hostname: config.hostname,
            username: config.username ? config.username : undefined,
            password: config.password ? config.password : undefined,
          };

          setConfig(cleanConfig);
        }
      )
    );
  });
}
