export class BadRequestError extends Error {
    _statusCode: string;
    message: string;
    constructor(statusCode: string, message: string) {
        super(message)
        this.message = message;
        this._statusCode = statusCode;
    }
}

export const ERR_USER_ALREADY_REGISTERED: BadRequestError = new BadRequestError("01", "user is already registered")
export const ERR_INVALID_ROLE: BadRequestError = new BadRequestError("02", "invalid role")