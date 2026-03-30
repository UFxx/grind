export interface InviteCode
{
	id         : string,
	code       : string,
	uses       : number,
	max_uses   : number,
	created_at : string
}

export interface FormattedInviteCode
{
	id        : string,
	code      : string,
	uses      : number,
	maxUses   : number,
	createdAt : string
}

export interface AddInviteCode
{
	code    : string | undefined,
	maxUses : string | undefined
}