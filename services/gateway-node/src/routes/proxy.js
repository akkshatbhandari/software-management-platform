import express from 'express';
import axios from 'axios';
import { authenticateToken } from '../middleware/auth.js';
import { requireRole } from '../middleware/roles.js';

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
            `${ENV.CORE_GO_BASE_URL}/projects`,
        );
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(502).json({
            error: "Failed to fetch projects from Core service",
        });
    }
});

router.post("/projects", authenticateToken, async(req, res)=>{
    try {
        const response = await axios.post(
            `${ENV.CORE_GO_BASE_URL}/projects`,
            req.body,
            {
                headers: {
                    "X-User-ID": req.user.user_id,
                }
            }
        );
        res.status(response.status).json(response.data);
    } catch (error) {
        if(error.response) {
            res.status(error.response.status).json(error.response.data);
        }else{
            res.status(502).json({
                error: "Core service is unreachable",
            });
        }
    }
});

router.get("/projects/all", authenticateToken, requireRole(["admin"]),
            async(req,res)=>{
                const response = await axios.get("http://localhost:3000/projects");
                res.json(response.data);
            }
);

export default router;
