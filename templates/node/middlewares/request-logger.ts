// @ts-nocheck
import {Request, Response, NextFunction} from "express";
import Logger from "@services/Logger.ts";

const RequestLogger = () => {
    return (req: Request, res: Response, next: NextFunction) => {
        Logger.LogImportant(`[${req.method}] - ${req.path}`);
        next();
    }
}

export default RequestLogger;