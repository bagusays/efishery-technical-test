import { BadRequestError } from './../error';
import { Response, Request, NextFunction } from 'express';

type errResponse = {
    code: string,
    message: string,
}

export function errorHandler(err: TypeError | BadRequestError, req: Request, res: Response, next: NextFunction): void {
    let resp: errResponse = {
        code: "-1",
        message: "fatal error! please contact the service owner"
    };
    if(err instanceof BadRequestError) {
        res.status(400)
        resp.code = err._statusCode;
        resp.message = err.message;
        res.json(resp)
        return
    }
    console.log("FATAL ERROR:", err.message) // should be contains all request meta data for tracing & debugging
    res.status(500)
    res.json(resp)
}