export interface User
{
	id          : string,
	name        : string,
	rank        : UserRank,
	level       : UserLevel,
	inviter?    : Inviter,
	avatar_url  : string,
	created_at  : string,
	telegram_id : number,
};
export interface FormattedUser
{
	id         : string,
	name       : string,
	rank       : UserRank,
	level      : FormattedUserLevel,
	inviter?   : FormattedInviter,
	avatarUrl  : string,
	createdAt  : string,
	telegramId : number
};

export interface UserResponse { data: User };

export interface UserRank
{
	id   : string,
	name : string
};

export interface UserLevel
{
	total_xp               : number,
	next_level             : number,
	current_level          : number,
	xp_to_next_level       : number,
	next_level_start_xp    : number,
	current_level_start_xp : number
};

export interface FormattedUserLevel
{
    totalXp             : number,
    nextLevel           : number,
    currentLevel        : number,
    xpToNextLevel       : number,
    nextLevelStartXp    : number,
    currentLevelStartXp : number
};

export type Inviter = Omit<User, 'inviter'>;

export type FormattedInviter = Omit<FormattedUser, 'inviter'>;