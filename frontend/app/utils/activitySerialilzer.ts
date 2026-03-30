import { type Activity, type FormattedActivity } from "~/types/activity";

export default (activities: Activity[]): FormattedActivity[] =>
{
	return activities.map(activity =>
		{
			return {
				id               : activity.id,
				rewards          : activityRewardsSerializer(activity.rewards),
				createdAt        : activity.created_at,
				description      : activity.description,
				activityCategory :
				{
					id   : activity.activity_category.id,
					name : activity.activity_category.name
				},
				tags             :
				[
					{
						name     : 'New',
						isActive : activity.is_new
					},
					{
						name     : 'Hard',
						isActive : activity.is_hard
					},
					{
						name     : 'Impact',
						isActive : activity.has_impact
					}
				]
			}
		}
	)
}