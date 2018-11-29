import { Command } from 'commander'

const program = new Command()

program
    .option('-c, --config <path>', 'configuration options file')
    .option('-H, --host <url> [port]', 'instance hostname or socket file')

program
    .command('contacts', 'find, list, or update contacts')

program.parse(process.argv)
