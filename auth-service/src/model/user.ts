import { ERR_INVALID_ROLE } from './../error';

export enum enumRole {
    ADMIN = "ADMIN",
    BASIC = "BASIC"
}

export function toRole(str: string): enumRole {
    switch(str) {
        case "ADMIN":
            return enumRole.ADMIN;
        case "BASIC":
            return enumRole.BASIC;
        default:
            throw ERR_INVALID_ROLE;
    }
}

export type user = {
    phone: string;
    name: string;
    role: enumRole;
    userName: string;
    password?: string;
    created_at?: string;
}