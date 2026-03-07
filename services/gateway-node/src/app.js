import express from 'express';
import cookieParser from 'cookie-parser';
import authRoutes from './routes/auth.js';
import {apiLimiter} from './middleware/rateLimit.js';
import {logger} from './middleware/logger.js';
import {requestId} from './middleware/requestId.js';


const app = express();

app.use(requestId);
app.use(logger);
app.use(express.json());
app.use(cookieParser());

app.use('/auth',authRoutes);

app.use("/api", apiLimiter);
export default app;