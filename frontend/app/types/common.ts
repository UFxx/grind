export interface SuccessResponse<T> { data: T };

export interface FetchError extends Error
{
	data?         : any,
	statusCode    : number,
	statusMessage : string
}