import jwt from 'jsonwebtoken';

function authenticateToken(req, res, next) {
    const authHeader = req.headers["authorization"];

    if (!authHeader) {
        return res.status(401).json({
            error: "Missing Authorization header"
        });
    }

    const token = authHeader.split(" ")[1];

    jwt.verify(token, process.env.JWT_SECRET, (err, decoded)=>{
        if(err) {
            return res.status(403).json({
                error: "Invalid or expired token"
            });
        }

        req.user = decoded;

        next();
    });
}

export {authenticateToken};