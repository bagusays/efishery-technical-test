import * as http from 'http';
import { Response } from 'express';

export interface IHttpServer {
    start(): http.Server
    stop(httpServer: http.Server): void;
}

export function responseJSON(res: Response, data: any, httpStatusCode: number = 200) {
    res.status(httpStatusCode)
    res.json(data)
}