export function IsLoggedIn(): boolean {
    return document.cookie.indexOf(process.env.SESSION_FLAG_COOKIE_NAME!) != -1;
}