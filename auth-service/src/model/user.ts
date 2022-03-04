export enum enumRole {
    ADMIN = "admin",
    BASIC = "basic"
}

export type user = {
    phone: string;
    name: string;
    role: enumRole;
    userName: string;
    password?: string;
    created_at?: string;
}