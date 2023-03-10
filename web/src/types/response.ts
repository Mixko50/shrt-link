export interface BasedResponse<T> {
    success: boolean
    data: T
}

export interface ShrtResponse {
    long_url: string
    slug: string
}

export interface ErrorResponse {
    error_message: string
    detail: string
}