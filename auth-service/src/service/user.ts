import { ERR_USER_ALREADY_REGISTERED, ERR_USER_NOT_FOUND  } from '../error';
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
        if(users.find(x => x.userName === args.userName && x.phone === args.phone)) {
            throw ERR_USER_ALREADY_REGISTERED;
        }

        args.password = generateRandomString(4)
        args.created_at = DateTime.now().toString()
        users.push(args);
        writeFileSync(this._fileName, JSON.stringify(users), "utf-8");
        
        return args.password;
    }

    findByPhone(phone: string): user {
        const usersjson = readFileSync(this._fileName, "utf-8");
        const users: user[] = JSON.parse(usersjson);
        const user: user = users.find(x => x.phone === phone);
        if (!user) {
            throw ERR_USER_NOT_FOUND
        }

        return user
    }
}