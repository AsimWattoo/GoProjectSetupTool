// @ts-nocheck
import {ErrorResponses, Errors} from "@/settings/errors-config.ts";
import ErrorResponse from "@/types/errors/ErrorResponse.ts";
import APIError from "@/types/errors/APIError.ts";
import ValidationResult from "@/types/generic/ValidationResult.ts";
import ApiResponse from "@/types/generic/ApiResponse.ts";

class ErrorService {

    private static _updateMessage(message: string, additionalData?: string[]) {
        let newMessage = message;
        if(additionalData) {
            for(let i = 0; i < additionalData.length; i++) {
                newMessage = newMessage.replace(`{${i + 1}}`, additionalData[i]);
            }
        }
        return newMessage;
    }

    // Returns list of all error codes
    static getAllErrorCodes() {
        return Object.values(Errors);
    }

    // Returns message for the specified code
    static findErrorMessages(codes: string[], fallback: string = 'Unidentified error has occurred'): ErrorResponse[] {
        const errorCodes = this.getAllErrorCodes();
        const responses = [] as ErrorResponse[];
        for(const code of codes) {
            if(errorCodes.includes(code)) {
                responses.push(ErrorResponses[code]);
            }
            else {
                responses.push(new ErrorResponse(Errors.UnidentifiedError, fallback));
            }
        }
        return responses;
    }

    static updateErrorMessages(errorResponses: ErrorResponse[], additionalData?: Record<string, string[]>): ErrorResponse[] {
        const updatedResponses = [] as ErrorResponse[];
        for(const response of errorResponses) {
            const responseData = additionalData ? additionalData[response.code] : undefined;
            if(responseData) {
                updatedResponses.push(new ErrorResponse(response.code, this._updateMessage(response.message, responseData)));
            }
            else {
                updatedResponses.push(response);
            }
        }
        return updatedResponses;
    }

    static generateAPIError(error: string, statusCode: number = 400, ...data: string[]) {
        const additionalData = {} as Record<string, string[]>;
        if(data && data.length > 0) {
            additionalData[error] = data;
        }
        return new APIError([error], statusCode, additionalData);
    }

    static generateAPIErrorFromResult(validation: ValidationResult): APIError {
        return new APIError(validation.errors, 400, validation.additionalData);
    }

    static generateAPIErrorFromApiResponse(response: ApiResponse) {
        const data = {} as Record<string, string[]>
        response.errors.forEach(e => {
            if(!data[e.code]) {
                data[e.code] = [] as string[]
            }
            data[e.code].push(e.message)
        })
        return new APIError(response.errors.map(e => e.code), 400, data);
    }

    static generateValidationResult(error: string, ...data: string[]): ValidationResult {
        const additionalData = {} as Record<string, string[]>;
        if(data && data.length > 0) {
            additionalData[error] = data;
        }
        return {
            errors: [error],
            additionalData: additionalData,
            success: false,
        };
    }
}

export default ErrorService;