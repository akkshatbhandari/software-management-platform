import express from 'express';
import cookieParser from 'cookie-parser';
import authRoutes from './routes/auth.js';
import {apiLimiter} from './middleware/rateLimit.js';


const app = express();

app.use(express.json());
app.use(cookieParser());

app.use('/auth',authRoutes);

app.use("/api", apiLimiter);
export default app;