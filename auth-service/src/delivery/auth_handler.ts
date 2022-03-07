import { ERR_TOKEN_IS_MISSING } from '../error';
import { user } from '../model/user';
import { Container } from '../container';
import { NextFunction, Response, Request, } from 'express';
import { ExpressRouteFunc } from './express'
import { responseJSON } from './common';

export class AuthHandler {
    login(container: Container): ExpressRouteFunc {
        type reqType = {
            phone: string;
            password: string;
        }

        type respType = {
            token: string;
        }

        return function(req: Request, res: Response, next?: NextFunction) {
            const userReq: reqType = req.body;
            
            const token = container.authService.login(userReq.phone, userReq.password);
            const resp: respType = {
                token: token,
            }

            responseJSON(res, resp);
        }
    }

    validate(container: Container): ExpressRouteFunc {
        type reqType = {
            phone: string;
            password: string;
        }

        return function(req: Request, res: Response, next?: NextFunction) {
            const authorizationHeader: string = req.headers.authorization;
            if (authorizationHeader === undefined){
                throw ERR_TOKEN_IS_MISSING;
            }
            const token: string = authorizationHeader.replace("Bearer ", "");
            if (token === "") {
                throw ERR_TOKEN_IS_MISSING;
            }
            
            const user: user = container.authService.validate(token);
            responseJSON(res, user);
        }
    }
}