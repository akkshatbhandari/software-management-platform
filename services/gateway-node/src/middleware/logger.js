function logger(req, res, next) {

    const startTime = Date.now();

    res.on("finish", ()=>{
        const duration = Date.now() - startTime;

        console.log({
            requestId: req.requestId,
            method: req.method,
            path: req.originalUrl,
            status: res.statusCode,
            duration: duration + "ms"
        });
    });

    next();
}

export {logger};