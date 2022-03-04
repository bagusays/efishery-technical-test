import { Container } from './../container';
import express, { Express, NextFunction, Response, Request, } from 'express';
import { ExpressRouteFunc } from './express'

export class AuthHandler {
    validate(container: Container): ExpressRouteFunc {
        return function(req: Request, res: Response, next?: NextFunction) {
            res.send(container.authService.generateJWT())
        }
    }
}