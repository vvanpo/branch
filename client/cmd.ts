import * as yargs from "yargs"

const args = yargs
    .option("c", {
        alias: "config",
        describe: "Path to configuration file",
        config: true,
    })
    .option("h", {
        alias: "host",
        describe: "Instance hostname",
        type: "string",
        default: "localhost",
    })
    .option("p", {
        alias: "port",
        describe: "Port number for server connection",
        type: "number",
        default: 8443,
    })

args.command({
    command: ["contacts", "co"],
    describe: "Find, list, or update contacts",
    builder: yargs => {
        const args = yargs.command({
            command: "find",
            describe: "Find a contact by e-mail",
        })

        args.command({
            command: "list",
            describe: "List contacts",
            builder: {
                filter: {
                    describe: "Filter contacts",
                },
            },
        })

        return args
    },
})

args.command("groups", "Enumerate or modify contact groups")

const argv = args.argv
