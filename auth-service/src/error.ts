export class BadRequestError extends Error {
    _statusCode: string;
    message: string;
    constructor(statusCode: string, message: string) {
        super(message)
        this.message = message;
        this._statusCode = statusCode;
    }
}

export const ERR_USER_ALREADY_REGISTERED: BadRequestError = new BadRequestError("01", "user with this phone & this username is already registered")
export const ERR_INVALID_ROLE: BadRequestError = new BadRequestError("02", "invalid role")
export const ERR_USER_NOT_FOUND: BadRequestError = new BadRequestError("03", "user not found")
export const ERR_INVALID_CREDENTIAL: BadRequestError = new BadRequestError("04", "invalid credential")
export const ERR_INVALID_TOKEN: BadRequestError = new BadRequestError("05", "invalid token")
export const ERR_TOKEN_IS_MISSING: BadRequestError = new BadRequestError("06", "token is missing")