export interface IUser
{
	id          : number,
	name        : string,
	rank        : IUserRank,
	level       : IUserLevel,
	inviter?    : TInviter,
	avatar_url  : string,
	created_at  : string,
	telegram_id : number,
};

export interface IUserRank
{
	id   : number,
	name : string
};

export interface IUserLevel
{
	total_xp               : number,
	next_level             : number,
	current_level          : number,
	xp_to_next_level       : number,
	next_level_start_xp    : number,
	current_level_start_xp : number
};

export type TInviter = Omit<IUser, 'inviter'>;