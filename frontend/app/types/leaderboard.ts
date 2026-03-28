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

export interface BaseEntry
{
	name: string,
	avatar_url: string,
};
export interface BaseFormattedEntry
{
	name: string,
	avatarURL: string
}

export interface Entry extends BaseEntry
{
	position: number,
	score: number
};
export interface FormattedEntry extends BaseFormattedEntry
{
	position: number,
	score: number
};

export interface MyEntry extends BaseEntry
{
	position: number | null,
	score: number | null,
	is_in_leaderboard: boolean
};
export interface FormattedMyEntry extends BaseFormattedEntry
{
	position: number | null,
	score: number | null,
	isInLeaderboard: boolean
};

export interface SeasonDetails
{
	season: Season,
	cta: string,
	entries: Entry[],
	me: MyEntry
};
export interface FormattedSeasonDetails
{
	season: FormattedSeason,
	cta: string,
	entries: FormattedEntry[],
	me: FormattedMyEntry
}