
function requireRole(allowedRoles = []) {
    return (req, res, next) => {
        if(!req.user || !req.user.role) {
            return res.status(403).
                        json({
                            error: "User role not found in token"
                        });
        }

        if(!allowedRoles.includes(req.user.role)) {
            return res.status(403).
                        json({
                            error: "Insufficient permissions to access this resource"
                        });
        }
        next();
    }
}

export {requireRole};