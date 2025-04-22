class ApiResponse {
    success: boolean = false;
    statusCode: number = 200;
    data?: Record<string, any> = {};
    errors: {code: string, message: string}[] = [];
}

export default ApiResponse;