import { getConfigUseCase } from "../../../application/get-config-use-case.js";

export function getConfigCommand(program) {
  program
    .command("config")
    .command("get")
    .action(() => {
      const config = getConfigUseCase();
      console.log(config);
    });
}
