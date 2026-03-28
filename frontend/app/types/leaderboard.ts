export interface Season
{
	id : string,
	name: string,
	period_start: string,
	period_end: string
};
export interface FormattedSeason
{
	id: string,
	name: string,
	periodStart: string,
	periodEnd: string
};

export interface Entry
{
	position: number,
	name: string,
	avatar_url: string,
	score: number
};
export interface FormattedEntry
{
	position: number,
	name: string,
	avatarURL: string,
	score: number
};

export interface MyEntry extends Entry
{
	is_in_leaderboard: boolean
};
export interface FormattedMyEntry extends FormattedEntry
{
	isInLeaderboard: boolean
};

export interface SeasonDetail
{
	season: Season,
	cta: string,
	entries: Entry[],
	me: MyEntry
};