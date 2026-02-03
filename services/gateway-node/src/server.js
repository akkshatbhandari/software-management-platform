import app from './app.js';
import {ENV} from './config/env.js';
import healthRoutes from './routes/health.js';
import proxyRoutes from './routes/proxy.js';

app.use("/",healthRoutes);
app.use("/api",proxyRoutes);

app.listen(ENV.PORT, () => {
    console.log(`Node.js Gateway is running on port ${ENV.PORT}`);
});