<script setup lang="ts">
	import type { FormattedEntry } from '~/types/leaderboard';

	const props = defineProps<{ entries: FormattedEntry[] }>();

	const [firstPlace, secondPlace, thirdPlace] = props.entries;
</script>

<template>
	<div class="leaderboard-podium">
		<div class="podium-item-wr">
			<LeaderboardPodiumItem
				v-if="secondPlace"
				:entry="secondPlace"
				v-motion-pop
				:delay="secondPlace.position * 100"
			/>
		</div>
		<div class="podium-item-wr">
			<LeaderboardPodiumItem
				v-if="firstPlace"
				:entry="firstPlace"
				:is-highlighted="true"
				v-motion-pop
				:delay="firstPlace.position * 100"
			/>
		</div>
		<div class="podium-item-wr">
			<LeaderboardPodiumItem
				v-if="thirdPlace"
				:entry="thirdPlace"
				v-motion-pop
				:delay="thirdPlace.position * 100"
			/>
		</div>
	</div>
</template>

<style lang="scss">
.leaderboard-podium
{
	width: 100%;

	display: flex;
	justify-content: space-around;
	align-items: flex-end;
}

.podium-item-wr { flex-basis: 33%; }
</style>