export interface BasedResponse {
	success: boolean;
	data: ShrtResponse;
	error: ErrorResponse;
}

export interface ShrtResponse {
	long_url: string;
	slug: string;
}

export interface ErrorResponse {
	error_message: string;
	detail: string;
}
