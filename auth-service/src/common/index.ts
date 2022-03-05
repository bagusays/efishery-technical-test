import { createHash } from 'crypto'

export function generateRandomString(len: number): string {
    const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]\{};:,./<>?";
    let text = "";
  
    for (let i = 0; i < len; i++) {
        text += charset.charAt(Math.floor(Math.random() * charset.length));
    }
  
    return text;
}

export function sha256(initialText: string, ...additionalText: string[]): string {
    let finalText: string = initialText;
    for (const d of additionalText) {
        finalText += "|" + d
    }
    return createHash("sha256").update(finalText).digest("hex")
}