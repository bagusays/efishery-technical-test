import { ERR_INVALID_CREDENTIAL, ERR_INVALID_TOKEN } from './../error';
import { UserService } from './user';
import { Config } from './../config/index';
import * as jwt from 'jsonwebtoken';
import { user } from '../model/user';
export class AuthService {
    _config: Config;
    _userService: UserService;

    constructor(cfg: Config, userService: UserService) {
        this._config = cfg;
        this._userService = userService;
    }

    login(phone: string, password: string): string {
        const expiresIn = 60 * 60; // an hour
        const user: user = this._userService.findByPhone(phone)
        if (user.password !== password) {
            throw ERR_INVALID_CREDENTIAL
        }

        const dataStoredInToken: user = {
            name: user.name,
            phone: user.phone,
            role: user.role,
            created_at: user.created_at,
            userName: user.userName
        };
        return jwt.sign(dataStoredInToken, this._config.JWT_SECRET, { expiresIn })
    }

    validate(token: string): user {
        try {
            return jwt.verify(token, this._config.JWT_SECRET) as user;
        } catch (error) {
            throw ERR_INVALID_TOKEN;
        }
    }
}