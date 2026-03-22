export type ToastType = 'success' | 'error' | 'info';

export interface Toast
{
	id: number,
	type: ToastType,
	text: string
}