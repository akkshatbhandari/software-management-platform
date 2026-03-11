import {Queue} from "bullmq";

import connection from 'redis.js';

const emailQueue = new Queue("emailQueue", {connection});

export {emailQueue};

