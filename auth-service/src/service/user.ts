import { Config } from './../config/index';
import { ERR_USER_ALREADY_REGISTERED, ERR_USER_NOT_FOUND  } from '../error';
import { user } from '../model/user';
import { writeFileSync, readFileSync, existsSync } from 'fs';
import { generateRandomString, sha256 } from '../common';
import { DateTime } from 'luxon'
export class UserService {
    _config: Config;

    constructor(cfg: Config) {
        this._config = cfg;
    }

    create(args: user): string {
        if (!existsSync(this._config.DB_FILENAME)) {
            writeFileSync(this._config.DB_FILENAME, "[]")
        }

        let usersjson = readFileSync(this._config.DB_FILENAME, "utf-8");
        let users: user[] = JSON.parse(usersjson);
        if(users.find(x => x.userName === args.userName && x.phone === args.phone)) {
            throw ERR_USER_ALREADY_REGISTERED;
        }

        const pw = generateRandomString(4)
        args.password = sha256(this._config.HASH_PEPPER, args.phone, pw)
        args.created_at = DateTime.now().toString()
        users.push(args);
        writeFileSync(this._config.DB_FILENAME, JSON.stringify(users), "utf-8");
        
        return pw;
    }

    findByPhone(phone: string): user {
        const usersjson = readFileSync(this._config.DB_FILENAME, "utf-8");
        const users: user[] = JSON.parse(usersjson);
        const user: user = users.find(x => x.phone === phone);
        if (!user) {
            throw ERR_USER_NOT_FOUND
        }

        return user
    }
}