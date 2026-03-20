export interface InviteCode
{
	id         : string,
	code       : string,
	uses       : number,
	max_uses   : number,
	created_at : string
}

export interface InviteCodeResponse { data: InviteCode[] };

export interface InviteCodeSuccessResponse { data: [] };

export interface FormattedInviteCode
{
	id        : string,
	code      : string,
	uses      : number,
	maxUses   : number,
	createdAt : string
}
