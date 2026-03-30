import { type ActivityReward, type FormattedActivityReward } from "~/types/activity";

export default (rewards: ActivityReward[]): FormattedActivityReward[] =>
{
	return rewards.map(reward =>
		{
			return {
				skillId   : reward.skill_id,
				xpAmount  : reward.xp_amount,
				skillName : reward.skill_name
			}
		}
	)
}