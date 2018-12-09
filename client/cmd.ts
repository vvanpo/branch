import * as yargs from "yargs"

const args = yargs
    .option("c", {
        alias: "config",
        config: true,
        describe: "Path to configuration file",
    })
    .option("h", {
        alias: "host",
        default: "localhost",
        describe: "Instance hostname",
        type: "string",
    })
    .option("p", {
        alias: "port",
        default: 8443,
        describe: "Port number for server connection",
        type: "number",
    })

args.command({
    builder: (yargs) => yargs
        .command({
            builder: {},
            command: "find <e-mail>",
            describe: "Find a contact by e-mail",
        })
        .command({
            builder: {
                filter: {
                    describe: "Filter contacts",
                },
            },
            command: "list",
            describe: "List contacts",
        }),
    command: ["contacts", "co"],
    describe: "Find, list, or update contacts",
})

args.command("groups", "Enumerate or modify contact groups")

const argv = args.argv
