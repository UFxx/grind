<script setup lang="ts">
	import type { FormattedEntry } from '~/types/leaderboard';

	const props = defineProps<
		{
			entry: FormattedEntry,
			isMe: boolean
		}
	>();
</script>

<template>
	<div
		class="leaderboard-entry"
		:class="{ 'leaderboard-entry--me': isMe }"
	>
		<div class="leaderboard-entry__main">
			<p class="leaderboard-entry__position">{{ entry.position }}</p>
			<p class="leaderboard-entry__name">{{ entry.name }}</p>
		</div>
		<p class="leaderboard-entry__score">{{ entry.score }}</p>
	</div>
</template>

<style lang="scss">
	.leaderboard-entry
	{
		padding: 10px 14px;
		border-left: 3px solid transparent;

		display: flex;
		justify-content: space-between;

		color: $gray;
		font-size: 12px;
		font-weight: 500;

		&__main
		{
			display: flex;
			gap: 20px;
		}

		&--me
		{
			color: $white;
			background-color: $darkGray;

			border-color: $white;
			border-radius: 0 4px 4px 0;
		}
	}
</style>