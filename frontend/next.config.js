/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    env: {
        API_URL: process.env.API_URL,
        REDIRECT_PARAM: process.env.REDIRECT_PARAM,
        CLIENT_RPC_API_TOKEN: process.env.CLIENT_RPC_API_TOKEN,
        SESSION_FLAG_COOKIE_NAME: process.env.SESSION_FLAG_COOKIE_NAME,
        DISCORD_INVITE_URL: process.env.DISCORD_INVITE_URL,
    },
    images: {
        domains: ["cdn.discordapp.com", "i.ytimg.com"]
    },
    experimental: {
        appDir: true,
    },
}

module.exports = nextConfig
