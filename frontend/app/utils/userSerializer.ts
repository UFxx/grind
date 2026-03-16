import { type User, type FormattedUser } from "~/types/user"

export default (user: User ): FormattedUser  =>
{
	return {
		id         : user.id,
		name       : user.name,
		rank       : user.rank,
		avatarUrl  : user.avatar_url,
		createdAt  : user.created_at,
		telegramId : user.telegram_id,
		inviter    : user?.inviter ? inviterSerializer(user.inviter) : undefined,
		level      :
		{
			totalXp             : user.level.total_xp,
			nextLevel           : user.level.next_level,
			currentLevel        : user.level.current_level,
			xpToNextLevel       : user.level.xp_to_next_level,
			nextLevelStartXp    : user.level.next_level_start_xp,
			currentLevelStartXp : user.level.current_level_start_xp,
		}
	}
}