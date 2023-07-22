/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    env: {
        API_URL: process.env.API_URL,
        CLIENT_RPC_API_TOKEN: process.env.CLIENT_RPC_API_TOKEN,
        SESSION_FLAG_COOKIE_NAME: process.env.SESSION_FLAG_COOKIE_NAME,
    },
    images: {
        domains: ["cdn.discordapp.com"]
    }
}

module.exports = nextConfig
