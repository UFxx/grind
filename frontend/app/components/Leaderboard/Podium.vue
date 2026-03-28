<script setup lang="ts">
	import type { FormattedEntry } from '~/types/leaderboard';

	const props = defineProps<{ entries: FormattedEntry[] }>();

	const [firstPlace, secondPlace, thirdPlace] = props.entries;
</script>

<template>
	<div class="leaderboard-podium">
		<LeaderboardPodiumItem
			v-if="firstPlace"
			class="leaderboard-podium__first"
			:entry="firstPlace"
			:is-highlighted="true"
			v-motion-pop
			:delay="firstPlace.position * 100"
		/>
		<LeaderboardPodiumItem
			v-if="secondPlace"
			class="leaderboard-podium__second"
			:entry="secondPlace"
			v-motion-pop
			:delay="secondPlace.position * 100"
		/>
		<LeaderboardPodiumItem
			v-if="thirdPlace"
			class="leaderboard-podium__third"
			:entry="thirdPlace"
			v-motion-pop
			:delay="thirdPlace.position * 100"
		/>
	</div>
</template>

<style lang="scss">
.leaderboard-podium
{
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	gap: 50px;
	align-items: end;

	&__first
	{
		grid-row: 1;
		grid-column: 2;
	}

	&__second
	{
		grid-row: 1;
		grid-column: 1;
	}

	&__third
	{
		grid-row: 1;
		grid-column: 3;
	}
}
</style>