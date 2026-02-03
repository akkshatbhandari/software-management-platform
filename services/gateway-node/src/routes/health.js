import express from 'express';

const Router = express.Router();

Router.get('/health', (req,res)=>{
    res.status(200).json({
        status:'UP', service:'gateway-node'
    })
});

export default Router;