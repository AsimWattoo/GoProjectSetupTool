// @ts-nocheck
const config = {
    PORT: process.env.PORT || 5001,
    ORIGINS: process.env.ORIGINS ? process.env.ORIGINS.split(',') : ['http://localhost:5173'],
}

export default config;