import { Response, Request, NextFunction } from 'express';

export function errorHandler(err: Error, req: Request, res: Response, next: NextFunction): void {
    res.status(500)
    res.send(err.message)
}