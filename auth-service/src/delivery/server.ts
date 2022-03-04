import * as http from 'http';

export interface IHttpServer {
    start(): http.Server
    stop(httpServer: http.Server): void;
}