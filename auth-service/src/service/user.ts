import { ERR_USER_ALREADY_REGISTERED  } from '../error';
import { user } from '../model/user';
import { writeFileSync, readFileSync, existsSync } from 'fs';
import { generateRandomString } from '../common';
import { DateTime } from 'luxon'
export class UserService {
    _fileName: string = "users.json"

    create(args: user): string {
        if (!existsSync(this._fileName)) {
            writeFileSync(this._fileName, "[]")
        }

        let usersjson = readFileSync(this._fileName, "utf-8");
        let users: user[] = JSON.parse(usersjson);
        if(users.find(x => x.userName === args.userName)) {
            throw ERR_USER_ALREADY_REGISTERED;
        }

        args.password = generateRandomString(4)
        args.created_at = DateTime.now().toString()
        users.push(args);
        writeFileSync(this._fileName, JSON.stringify(users), "utf-8");
        
        return args.password;
    }
}