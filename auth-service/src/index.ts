import { initContainer } from './container';
import { ExpressHTTP } from './delivery/express';
import * as dotenv from 'dotenv'
import { Config } from './config';
import { writeFileSync, existsSync } from 'fs';

try {
    const config: Config = dotenv.config({ path: 'config/config.env' }).parsed;

    if (!existsSync(config.DB_FILENAME)) {
        writeFileSync(config.DB_FILENAME, "[]")
    }

    const server = new ExpressHTTP(initContainer(config))
    const http = server.start()

    process.on('SIGTERM', server.stop(http));
    process.on('SIGINT', server.stop(http));
} catch (e) {
    console.log(e);
}
