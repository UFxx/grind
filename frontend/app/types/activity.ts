import { type Pagination } from "~/types/pagination"

export interface ActivitiesResponse
{
	pagination: Pagination,
	items: Activity[]
};

export interface Activity
{
	id                : string,
	is_new            : boolean,
	rewards           : ActivityReward[],
	is_hard           : boolean,
	created_at        : string
	has_impact        : boolean,
	description       : string,
	activity_category : ActivityCategory
};

export interface FormattedActivity
{
	id               : string,
	tags             : FormattedActivityTag[]
	rewards          : FormattedActivityReward[],
	createdAt        : string,
	description      : string,
	activityCategory : ActivityCategory
};

export interface FormattedActivityTag
{
	name     : 'New' | 'Hard' | 'Impact',
	isActive : boolean
}

interface ActivityCategory
{
	id   : string,
	name : string
};

export interface ActivityReward
{
	skill_id   : string,
	xp_amount  : number
	skill_name : string,
};

export interface FormattedActivityReward
{
	skillId   : string,
	xpAmount  : number,
	skillName : string
};

export interface AddActivity
{
	mode        : string,
	description : string
}