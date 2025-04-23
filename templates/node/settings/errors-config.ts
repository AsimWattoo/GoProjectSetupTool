// @ts-nocheck
import ErrorResponse from "@/types/errors/ErrorResponse.ts";

const Errors = {
    NotFound: 'not_found',
    UnidentifiedError: "unidentified_error",
    Unauthorized: 'unauthorized',
    Forbidden: 'forbidden_access',
    InvalidCredentials: 'invalid_credentials',
    RateLimitingReached: "rate_limit_reached",
    AlreadyInUse: "already_in_use",
    ModelNotRegistered: "model_not_registered",
}

const ErrorResponses = {
    [Errors.UnidentifiedError]: new ErrorResponse("unidentified_error", "An unidentified error occurred."),
    [Errors.NotFound]: new ErrorResponse(Errors.NotFound, "{1} not found"),
    [Errors.Unauthorized]: new ErrorResponse(Errors.Unauthorized, 'Unauthorized access'),
    [Errors.Forbidden]: new ErrorResponse(Errors.Forbidden, 'Forbidden access to this resource'),
    [Errors.InvalidCredentials]: new ErrorResponse(Errors.InvalidCredentials, "Invalid credentials"),
    [Errors.RateLimitingReached]: new ErrorResponse(Errors.RateLimitingReached, 'You have exceeded the maximum number of requests allowed per minute. Please try again later'),
    [Errors.AlreadyInUse]: new ErrorResponse(Errors.AlreadyInUse, `The {1} is already in use. Cannot use again`),
    [Errors.ModelNotRegistered]: new ErrorResponse(Errors.ModelNotRegistered, "The model {1} is not registered with the Database Provider"),
}


export {
    Errors,
    ErrorResponses
};