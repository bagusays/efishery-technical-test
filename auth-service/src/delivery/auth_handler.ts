import { ERR_INVALID_ROLE } from './../error';
import { enumRole, registry } from './../model/auth';
import { Container } from './../container';
import express, { Express, NextFunction, Response, Request, } from 'express';
import { ExpressRouteFunc } from './express'
import { responseJSON } from './common';

export class AuthHandler {
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
            const item: reqType = req.body;
            
            if (!(item.role in enumRole)) {
                throw ERR_INVALID_ROLE;
            }

            const arg: registry = {
                name: item.name,
                phone: item.phone,
                role: item.role,
                userName: item.userName
            }
            
            const password = container.authService.create(arg)
            const resp: respType = {
                password: password,
            }

            responseJSON(res, resp)
        }
    }
}