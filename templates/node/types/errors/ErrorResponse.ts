class ErrorResponse {
    message: string = '';
    code: string = '';

    constructor(code: string, message: string) {
        this.code = code;
        this.message = message;
    }

    getError() {
        return {
            code: this.code,
            message: this.message
        }
    }
}

export default ErrorResponse;