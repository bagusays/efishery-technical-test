import { Config } from './config/index';
import { Auth } from "./service/auth"

export type Container = {
    config: Config;
    authService: Auth;
}

export function initContainer(cfg: Config): Container {
    return {
    authService: new Auth(),
        config: cfg
    }
}