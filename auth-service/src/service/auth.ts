import { ERR_USER_ALREADY_REGISTERED  } from './../error';
import { registry } from './../model/auth';
import { writeFileSync, readFileSync, existsSync } from 'fs';
import { generateRandomString } from '../common';
import { DateTime } from 'luxon'
export class Auth {
    _fileName: string = "users.json"

    create(args: registry): string {
        if (!existsSync(this._fileName)) {
            writeFileSync(this._fileName, "[]")
        }

        let usersjson = readFileSync(this._fileName, "utf-8");
        let users: registry[] = JSON.parse(usersjson);
        if(users.find(x => x.userName === args.userName)) {
            throw ERR_USER_ALREADY_REGISTERED;
        }

        args.password = generateRandomString(4)
        args.timestamp = DateTime.now().toString()
        users.push(args);
        writeFileSync(this._fileName, JSON.stringify(users), "utf-8");
        
        return args.password;
    }
}