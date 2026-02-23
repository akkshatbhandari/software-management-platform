import express from 'express';
import authRoutes from './routes/auth.js';
import {apiLimiter} from './middleware/rateLimit.js';


const app = express();

app.use(express.json());

app.use('/auth',authRoutes);

app.use("/api", apiLimiter);
export default app;