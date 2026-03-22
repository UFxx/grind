export interface UiInput
{
	type?        : 'text' | 'number',
	inputmode?   : 'text' | 'numeric',
	placeholder : string,
	onlyNumbers? : boolean
}