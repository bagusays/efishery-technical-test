import { AuthHandler } from './auth_handler';
import { Container } from './../container';
import { IHttpServer } from './common';
import * as http from 'http';
import express, { Express, NextFunction, Response, Request, } from 'express';
import { UserHandler } from './user_handler'
import { errorHandler } from './error_handler';

export type ExpressRouteFunc = (req: Request, res: Response, next?: NextFunction) => void | Promise<void>;

export class ExpressHTTP implements IHttpServer {
    _port: number = 8080
    _userHandler: UserHandler;
    _authHandler: AuthHandler;
    _httpServer: Express;
    _container: Container;

    constructor(container: Container) {
        this._httpServer = express();
        this._httpServer.use(express.json())
        this._container = container;

        this._authHandler = new AuthHandler();
        this._userHandler = new UserHandler();

        this.registerRoutes();

        this._httpServer.use(errorHandler);
    }

    registerRoutes() {
        this._httpServer.post("/api/user", this._userHandler.create(this._container))
        this._httpServer.post("/api/auth/login", this._authHandler.login(this._container))
        this._httpServer.get("/api/auth/validate", this._authHandler.validate(this._container))
    }

    start(): http.Server {
        return this._httpServer.listen(this._port, () => {
            console.log(`⚡️[server]: Server is running at https://localhost:${this._port}`);
        });
    }

    stop(http: http.Server) {
        return function() {
            console.log('SIGTERM signal received: closing HTTP server')
            http.close(() => {
                console.log('HTTP server closed')
                process.exit(0);
            })
        }
    }
}