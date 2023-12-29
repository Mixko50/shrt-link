export interface BasedResponse {
	success: boolean;
	payload: ShrtResponse;
	message: string;
}

export interface ShrtResponse {
	original_url: string;
	slug: string;
}