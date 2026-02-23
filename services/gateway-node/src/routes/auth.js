import express from 'express';

import jwt from 'jsonwebtoken';

import axios from 'axios';

import {authLimiter} from '../middleware/rateLimit.js';


const router = express.Router();

router.post("/login",authLimiter,async(req,res)=>{
    const { email, password } = req.body;

    try {
        const response = await axios.post(
            "http://localhost:3000/auth/login",
            req.body
        );

        const {id, email, role} = response.data;

        const accessToken = jwt.sign(
            {
                user_id: id,
                email,
                role,
            },
            process.env.JWT_SECRET,
            {
                expiresIn: process.env.JWT_EXPIRES_IN
            }
        );

        const refreshToken = jwt.sign(
            {
                user_id:id
            },
            process.env.JWT_REFRESH_SECRET,
            {
                expiresIn: process.env.REFRESH_TOKEN_EXPIRES_IN
            }
        )

        await axios.post("http://localhost:3000/auth/refresh/store", {
            user_id: id,
            token: refreshToken
        });
        
        res.json({ accessToken, refreshToken });

    } catch (error) {
        res.status(401).json({
            error: "Invalid email or password"
        });
    }
});

router.post("/refresh", async(req, res)=>{
    const {refreshToken} = req.body;

    if(!refreshToken) {
        return res.status(401).json({
            error: "Refresh token is required"
        });
    }

    try {
        const payload = jwt.verify(refreshToken, process.env.JWT_SECRET);

        const newAccessToken = jwt.sign(
            {
                user_id: payload.user_id,
            },
            process.env.JWT_SECRET,
            {
                expiresIn: process.env.JWT_EXPIRES_IN
            }
        )

        res.json({ accessToken: newAccessToken });
    } catch (error) {
        return res.status(403).json({
            error: "Invalid refresh token"
        })
    }
});

export default router;