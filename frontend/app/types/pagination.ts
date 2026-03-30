export interface Pagination
{
	limit        : number,
	next_page    : number,
	total_pages  : number,
	current_page : number
};

export interface FormattedPagination
{
	limit       : number,
	nextPage    : number,
	totalPages  : number,
	currentPage : number
}