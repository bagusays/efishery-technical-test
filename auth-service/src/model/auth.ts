export enum enumRole {
    ADMIN = "admin",
    USER = "user"
}

export type registry = {
    phone: string;
    name: string;
    role: enumRole;
    userName: string;
    password?: string;
    timestamp?: string;
}