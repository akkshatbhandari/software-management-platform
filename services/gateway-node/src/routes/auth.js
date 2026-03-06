import express from 'express';

import jwt from 'jsonwebtoken';

import axios from 'axios';

import {authLimiter} from '../middleware/rateLimit.js';
import {ENV} from '../config/env.js';


const router = express.Router();

const REFRESH_COOKIE_NAME = 'refresh_token';

function refreshCookieOptions() {
    const isProd = process.env.NODE_ENV === "production";
    return {
        httpOnly: true,
        secure: isProd,
        sameSite: "strict",
        path: "/auth",
    };
}

router.post("/login",authLimiter,async(req,res)=>{
    const { email, password } = req.body;

    try {
        const response = await axios.post(
            `${ENV.CORE_GO_BASE_URL}/auth/login`,
            req.body
        );

        const {id, email, role} = response.data;

        const accessToken = jwt.sign(
            {
                token_type: "access",
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
                token_type: "refresh",
                user_id:id
            },
            process.env.JWT_REFRESH_SECRET,
            {
                expiresIn: process.env.REFRESH_TOKEN_EXPIRES_IN
            }
        )

        await axios.post(`${ENV.CORE_GO_BASE_URL}/auth/refresh/store`, {
            user_id: id,
            token: refreshToken
        });
        
        res.cookie(REFRESH_COOKIE_NAME, refreshToken, refreshCookieOptions());
        res.json({ accessToken });

    } catch (error) {
        res.status(401).json({
            error: "Invalid email or password"
        });
    }
});

router.post("/refresh", authLimiter, async(req, res)=>{
    const refreshToken = req.cookies?.[REFRESH_COOKIE_NAME];

    if(!refreshToken) {
        return res.status(401).json({
            error: "Refresh token is required"
        });
    }

    try {
        const payload = jwt.verify(refreshToken, process.env.JWT_REFRESH_SECRET);

        await axios.post(`${ENV.CORE_GO_BASE_URL}/auth/refresh/validate`, {
            user_id: payload.user_id,
            token: refreshToken,
        });

        const newRefreshToken = jwt.sign(
            {
                token_type: "refresh",
                user_id: payload.user_id,
            },
            process.env.JWT_REFRESH_SECRET,
            {
                expiresIn: process.env.REFRESH_TOKEN_EXPIRES_IN
            }
        );

        await axios.post(`${ENV.CORE_GO_BASE_URL}/auth/refresh/store`, {
            user_id: payload.user_id,
            token: newRefreshToken,
        });

        await axios.post(`${ENV.CORE_GO_BASE_URL}/auth/refresh/revoke`, {
            user_id: payload.user_id,
            token: refreshToken,
        });

        const newAccessToken = jwt.sign(
            {
                token_type: "access",
                user_id: payload.user_id,
            },
            process.env.JWT_SECRET,
            {
                expiresIn: process.env.JWT_EXPIRES_IN
            }
        )

        res.cookie(REFRESH_COOKIE_NAME, newRefreshToken, refreshCookieOptions());
        res.json({ accessToken: newAccessToken });
    } catch (error) {
        return res.status(403).json({
            error: "Invalid refresh token"
        })
    }
});

router.post("/logout", authLimiter, async(req, res)=>{
    const refreshToken = req.cookies?.[REFRESH_COOKIE_NAME];

    if(!refreshToken) {
        // Logout should be idempotent.
        res.clearCookie(REFRESH_COOKIE_NAME, refreshCookieOptions());
        return res.status(204).send();
    }

    try {
        const payload = jwt.verify(refreshToken, process.env.JWT_REFRESH_SECRET);

        await axios.post(`${ENV.CORE_GO_BASE_URL}/auth/refresh/revoke`, {
            user_id: payload.user_id,
            token: refreshToken,
        });

        res.clearCookie(REFRESH_COOKIE_NAME, refreshCookieOptions());
        return res.status(204).send();
    } catch (error) {
        res.clearCookie(REFRESH_COOKIE_NAME, refreshCookieOptions());
        return res.status(204).send();
    }
});

export default router;