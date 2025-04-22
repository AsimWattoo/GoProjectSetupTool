// @ts-nocheck
import {NextFunction, Request, Response} from 'express'
import APIError from "@/types/errors/APIError.ts";
import ErrorService from "@services/ErrorService";
import Logger from "@services/Logger.ts";

const ErrorHandler = () => {
    return async (error: Error, req: Request, res: Response, next: NextFunction) => {

        const apiError = error as APIError;
        if(apiError.errorCodes) {
            let errorResponses = ErrorService.findErrorMessages(apiError.errorCodes, error.message);
            errorResponses = ErrorService.updateErrorMessages(errorResponses, apiError.additionalData);
            const statusCode = apiError.statusCode || 500;
            const errors = errorResponses.map(res => res.getError());
            let response = {errors: errors};
            if(apiError.responseContent) {
                response = {...response, ...apiError.responseContent};
            }
            res.status(statusCode).json(response);
            Logger.LogError(`[${statusCode}] - ${errorResponses.map(res => `[${res.code} => ${res.message}]`)}`);
        }
        else {
            let response = {errors: [{message: error.message}]};
            if(apiError.responseContent) {
                response = {...response};
            }
            res.status(500).json(response);
            Logger.LogError(`[500] - [${error.message}]`);
        }
        return;
    }
}

export default ErrorHandler;