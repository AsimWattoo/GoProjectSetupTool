class APIError extends Error {
    errorCodes: string[];
    additionalData?: Record<string, string[]>;
    statusCode: number = 400;
    responseContent?: Record<string, any>;

    constructor(errorCodes: string[], statusCode: number = 400, additionalData?: Record<string, string[]>, responseContent?: Record<string, any>) {
        super();
        this.statusCode = statusCode;
        this.errorCodes = errorCodes;
        this.additionalData = additionalData;
        this.responseContent = responseContent;
    }
}

export default APIError;