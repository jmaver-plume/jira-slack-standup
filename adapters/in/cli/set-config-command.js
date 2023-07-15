import prompt from "prompt";

export function setConfigCommand(program) {
  program
    .command("config")
    .command("set")
    .action(async () => {
      prompt.start();
      const { username, password, hostname } = await new Promise(
        (resolve, reject) =>
          prompt.get(
            {
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
            },
            (err, data) => {
              if (err) {
                return reject(err);
              }
              return resolve(data);
            }
          )
      );

      setConfigUseCase(username, password, hostname);
    });
}
