export interface Skill
{
	id                     : string,
	name                   : string,
	total_xp               : number,
	next_level             : number,
	current_level          : number,
	current_level_start_xp : number,
	next_level_start_xp    : number,
	xp_to_next_level       : number,
	items?                 : Subskill[]
};

export interface FormattedSkill
{
	id                  : string,
	name                : string,
	totalXp             : number,
	nextLevel           : number,
	currentLevel        : number,
	xpToNextLevel       : number,
	nextLevelStartXp    : number,
	currentLevelStartXp : number,
	subskills?          : FormattedSubskill[]
};

export type Subskill = Omit<Skill, 'items'>;

export type FormattedSubskill = Omit<FormattedSkill, 'items'>