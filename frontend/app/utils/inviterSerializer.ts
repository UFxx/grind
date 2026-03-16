import { type Inviter, type FormattedInviter } from "~/types/user"

export default (inviter: Inviter): FormattedInviter =>
{
	return {
		id         : inviter.id,
		name       : inviter.name,
		rank       : inviter.rank,
		avatarUrl  : inviter.avatar_url,
		createdAt  : inviter.created_at,
		telegramId : inviter.telegram_id,
		level      : {
			totalXp             : inviter.level.total_xp,
			nextLevel           : inviter.level.total_xp,
			currentLevel        : inviter.level.current_level,
			xpToNextLevel       : inviter.level.xp_to_next_level,
			nextLevelStartXp    : inviter.level.next_level_start_xp,
			currentLevelStartXp : inviter.level.current_level_start_xp
		},
	}
}