import { ERR_INVALID_ROLE } from '../error';
import { enumRole, user } from '../model/user';
import { Container } from '../container';
import { NextFunction, Response, Request, } from 'express';
import { ExpressRouteFunc } from './express'
import { responseJSON } from './common';

export class UserHandler {
    create(container: Container): ExpressRouteFunc {
        type reqType = {
            phone: string;
            name: string;
            role: enumRole;
            userName: string;
        }

        type respType = {
            password: string;
        }

        return function(req: Request, res: Response, next?: NextFunction) {
            const userReq: reqType = req.body;
            
            if (!(userReq.role in enumRole)) {
                throw ERR_INVALID_ROLE;
            }

            const arg: user = {
                name: userReq.name,
                phone: userReq.phone,
                role: userReq.role,
                userName: userReq.userName
            }
            
            const password = container.userService.create(arg)
            const resp: respType = {
                password: password,
            }

            responseJSON(res, resp)
        }
    }
}