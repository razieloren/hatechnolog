/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    env: {
        API_WS_URL: process.env.API_WS_URL,
        MAIN_REPO_URL: process.env.MAIN_REPO_URL,
        MAIN_REPO_COMMITS_URL: process.env.MAIN_REPO_COMMITS_URL,
        MAIN_REPO_CONTRIBUTORS_URL: process.env.MAIN_REPO_CONTRIBUTORS_URL,
        DISCORD_INVITE_URL: process.env.DISCORD_INVITE_URL,
        YOUTUBE_CHANNEL_URL: process.env.YOUTUBE_CHANNEL_URL,
    }
}

module.exports = nextConfig
