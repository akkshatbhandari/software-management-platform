import express from 'express';
import axios from 'axios';

const router = express.Router();

import {ENV} from '../config/env.js';

router.get('/health',async(req,res)=>{
    try {
        const response = await axios.get(`${ENV.CORE_GO_BASE_URL}/health`);
        res.status(response.status).json({
            gateway: "OK",
            coreService: response.data
        });
    } catch (error) {
        res.status(502).json({
            error: "Core Go service is unreachable",
        });
    }
});

router.get('/projects', async(req, res)=>{
    try {
        const response = await axios.get(
            `${ENV.CORE_GO_BASE_URL}/projects`
        );
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(502).json({
            error: "Failed to fetch projects from Core service",
        });
    }
});

export default router;
