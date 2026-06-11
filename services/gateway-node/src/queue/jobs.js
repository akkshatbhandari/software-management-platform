import {Queue} from "bullmq";

import connection from '../queue/redis.js';

const emailQueue = new Queue("emailQueue", {connection});

export {emailQueue};

