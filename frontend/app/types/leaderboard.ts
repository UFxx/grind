export interface Season
{
	id           : string,
	name         : string,
	period_end   : string,
	period_start : string
};
export interface FormattedSeason
{
	id          : string,
	name        : string,
	periodEnd   : string,
	periodStart : string
};

export interface BaseEntry
{
	name       : string,
	avatar_url : string,
};
export interface BaseFormattedEntry
{
	name      : string,
	avatarURL : string
}

export interface Entry extends BaseEntry
{
	score    : number,
	position : number
};
export interface FormattedEntry extends BaseFormattedEntry
{
	score    : number,
	position : number
};

export interface MyEntry extends BaseEntry
{
	score             : number | null,
	position          : number | null,
	is_in_leaderboard : boolean
};
export interface FormattedMyEntry extends BaseFormattedEntry
{
	score           : number | null,
	position        : number | null,
	isInLeaderboard : boolean
};

export interface SeasonDetails
{
	me      : MyEntry,
	cta     : string,
	season  : Season,
	entries : Entry[]
};
export interface FormattedSeasonDetails
{
	me      : FormattedMyEntry,
	cta     : string,
	season  : FormattedSeason,
	entries : FormattedEntry[]
}