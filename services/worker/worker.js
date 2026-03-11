import {Worker} from "bullmq";

import connection from '../gateway-node/src/queue/redis.js';

const worker = new Worker(
    "emailQueue",
    async job =>{

        console.log("Processing job: ", job.name);

        if(job.name ==="sendWelcomeEmail") {

            const { email } = job.data;

            console.log("Sending welcome email to:", email);

            await new Promise(r => setTimeout(r, 2000));

            console.log("Email sent");
        }
    },
    {connection}
);