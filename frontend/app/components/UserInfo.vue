<script setup lang="ts">
	import { type FormattedUser } from '~/types/user';

	defineProps<{ user: FormattedUser }>();
</script>

<template>
	<div class="user-info">
		<UserAvatar
			:image="user.avatarUrl"
			:level-percent="getLeverPercent(user.level)"
			:level="user.level.currentLevel"
		/>
		<div class="user-info__text">
			<p class="user-info__tag">@{{ user.name }}</p>
			<p class="user-info__rank">{{ user.rank.name }}</p>
		</div>
	</div>
</template>

<style lang='scss' scoped>
	.user-info
	{
		row-gap: 10px;

		display: flex;
		align-items: center;
		flex-direction: column;
	}

	.user-info__text
	{
		row-gap: 2px;

		display: flex;
		align-items: center;
		flex-direction: column;
	}

	.user-info__tag { line-height: 19px; }

	.user-info__rank
	{
		color: $gray;
		font-size: 12px;
		line-height: 16px;
	}
</style>