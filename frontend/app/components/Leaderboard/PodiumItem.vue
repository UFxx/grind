<script setup lang="ts">
	import type { FormattedEntry } from '~/types/leaderboard';

	const props = defineProps<
		{
			entry: FormattedEntry,
			isHighlighted?: boolean
		}
	>();
</script>

<template>
	<div
		class="leaderboard-podium-item"
		:class="{ 'leaderboard-podium-item--highlighted': isHighlighted }"
	>
		<div class="leaderboard-podium-item__avatar-wr">
			<img
				class="leaderboard-podium-item__avatar"
				:src="entry.avatarURL"
				alt="Avatar"
			/>
			<p class="leaderboard-podium-item__position">{{ entry.position }}</p>
		</div>
		<p class="leaderboard-podium-item__name">{{ entry.name }}</p>
		<p class="leaderboard-podium-item__score">{{ entry.score }}</p>
	</div>
</template>

<style lang="scss">
.leaderboard-podium-item
{
	display: flex;
	flex-direction: column;
	align-items: center;
	gap: 2px;

	font-weight: 700;
	font-size: 12px;

	&__avatar-wr
	{
		width: 60px;
		height: 60px;

		position: relative;

		img
		{
			width: 100%;
			height: 100%;
			object-fit: cover;
			border-radius: 100%;
		}
	}

	&__position
	{
		position: absolute;
		bottom: 0;
		right: -6px;

		background-color: $darkGray;

		border-radius: 100%;
		width: 20px;
		height: 20px;

		display: flex;
		align-items: center;
		justify-content: center;
	}

	&__score
	{
		font-size: 10px;
		font-weight: 500;
		color: $gray;
	}
}

.leaderboard-podium-item--highlighted
{
	.leaderboard-podium-item__avatar-wr
	{
		width: 70px;
		height: 70px;
	}

	.leaderboard-podium-item__position
	{
		width: 28px;
		height: 28px;
		font-size: 14px;

		background-color: $white;
		color: $black;
	}

	.leaderboard-podium-item__name,
	.leaderboard-podium-item__score
	{
		font-weight: 700;
		color: $white;
	}

	.leaderboard-podium-item__name { font-size: 14px; }
	.leaderboard-podium-item__score { font-size: 12px; }
}
</style>