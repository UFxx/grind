<script setup lang="ts">
	import type { FormattedSeason } from '~/types/leaderboard';

	const props = defineProps<{ season: FormattedSeason }>();

	const seasonPeriodEnd = new Date(props.season.periodEnd);
</script>

<template>
	<div class="leaderboard-season">
		<p class="leaderboard-season__name">{{ season.name }}</p>
		<Countdown
			class="leaderboard-season__countdown"
			:date="seasonPeriodEnd"
			v-slot="{ days, hours, minutes, seconds}"
		>
			TIME REMAINING: {{ days }}D {{ hours }}H {{ minutes }}M {{ seconds }}S
		</Countdown>
	</div>
</template>

<style lang="scss">
	.leaderboard-season
	{
		text-align: center;

		&__name
		{
			font-size: 26px;
			font-weight: 700;
		}

		&__countdown
		{
			color: $gray;
			font-size: 12px;
			font-weight: 400;
		}
	}
</style>