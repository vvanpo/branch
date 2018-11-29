import * as yargs from "yargs"

const args = yargs
    .option("c", {
        alias: "config",
        describe: "Path to configuration file",
    })
    .option("h", {
        alias: "host",
        describe: "Instance hostname or socket file",
    })
    .command("contacts", "Find, list, or update contacts")

const argv = args.argv
