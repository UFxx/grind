<script setup lang="ts">
	import { type IUser } from '~/types/user';

	const props = defineProps<{user: IUser}>();

	const levelPercent = computed(() =>
		{
			const { total_xp, current_level_start_xp, next_level_start_xp } = props.user.level;

			if (next_level_start_xp <= current_level_start_xp) return 100;
			if (total_xp >= next_level_start_xp) return 100;
			if (total_xp <= current_level_start_xp) return 0;

			const progress = total_xp - current_level_start_xp;
			const levelRange = next_level_start_xp - current_level_start_xp;

			return Math.ceil((progress / levelRange) * 100);
		}
	)
</script>

<template>
	<div class="user-info">
		<UserAvatar
			:image="user.avatar_url"
			:level-percent
			:level="user.level.current_level"
		/>
		<p>{{ user.name }}</p>
		<p>{{ user.rank }}</p>
	</div>
</template>

<style lang='scss' scoped>
	.user-info
	{
		display: flex;
		align-items: center;
		flex-direction: column;

		row-gap: 10px;
	}
</style>