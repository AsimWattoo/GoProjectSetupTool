class ValidationResult {
    success: boolean = false;
    errors: string[] = [];
    data?: any;
    additionalData: Record<string, string[]> = {};

    constructor(success: boolean, errors: string[] = [], additionalData: Record<string, string[]> = {}) {
        this.success = success;
        this.errors = errors;
        this.additionalData = additionalData;
    }
}

export default ValidationResult;