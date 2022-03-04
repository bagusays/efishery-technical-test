import { Config } from './config/index';
import { UserService } from "./service/user"

export type Container = {
    config: Config;
    userService: UserService;
}

export function initContainer(cfg: Config): Container {
    return {
    userService: new UserService(),
        config: cfg
    }
}