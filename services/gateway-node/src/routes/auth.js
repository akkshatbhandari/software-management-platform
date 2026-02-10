import express from 'express';

import jwt from 'jsonwebtoken';

import axios from 'axios';


const router = express.Router();

router.post("/login",async(req,res)=>{
    const { email, password } = req.body;

    try {
        const response = await axios.post(
            "http://localhost:3000/auth/login",
            {email, password}
        );

        const {id, email: userEmail} = response.data;

        const token = jwt.sign(
            {
                user_id: id,
                email: userEmail
            },
            process.env.JWT_SECRET,
            {
                expiresIn: process.env.JWT_EXPIRES_IN
            }
        );
        
        res.json({ token });

    } catch (error) {
        res.status(401).json({
            error: "Invalid email or password"
        });
    }
});

export default router;