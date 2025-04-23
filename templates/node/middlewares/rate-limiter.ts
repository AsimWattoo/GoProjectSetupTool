// @ts-nocheck
import {rateLimit} from "express-rate-limit"
import {NextFunction, Response, Request} from "express";
import APIError from "@/types/errors/APIError.ts";
import {Errors} from "@/settings/errors-config.ts";

const RateLimiter = () => {
    return rateLimit({
        windowMs: 1000 * 60,
        limit: 10000,
        handler: (req: Request, response: Response, next: NextFunction) => {
            throw new APIError([Errors.RateLimitingReached])
        },
    })
}

export default RateLimiter;