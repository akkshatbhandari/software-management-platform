import {v4 as uuidv4} from 'uuid';

function requestId(req, res, next) {
    const id = uuidv4();

    req.requestId = id;

    res.setHeader('X-Request-ID', id);

    next();
}

export {requestId};