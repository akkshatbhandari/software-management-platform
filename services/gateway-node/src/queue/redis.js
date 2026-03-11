import {Redis} from 'ioredis';

const connection = new Redis({
    host: process.env.REDIS_HOST,
    port: process.env.REDIS_PORT,
});

export default connection;

