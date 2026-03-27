export interface UiInput
{
	type?        : 'text' | 'number',
	inputmode?   : 'text' | 'numeric',
	fullWidth?   : boolean,
	placeholder  : string,
	onlyNumbers? : boolean
}