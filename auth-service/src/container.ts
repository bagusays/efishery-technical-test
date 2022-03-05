import { AuthService } from './service/auth';
import { Config } from './config/index';
import { UserService } from "./service/user"

export type Container = {
    config: Config;
    authService: AuthService;
    userService: UserService;
}

export function initContainer(cfg: Config): Container {
    const userService = new UserService();
    return {
        authService: new AuthService(cfg, userService),
        userService: userService,
        config: cfg
    }
}