import dotenv from 'dotenv';

dotenv.config();

const ENV = {
    PORT: process.env.PORT || 3000,
    CORE_GO_BASE_URL: process.env.CORE_GO_BASE_URL || 'http://localhost:3000',
};

export {ENV};