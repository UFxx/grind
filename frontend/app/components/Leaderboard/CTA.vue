<script setup lang="ts">
	const props = defineProps(
		{
			text:
			{
				type     : String,
				required : true
			}
		}
	);
</script>

<template>
	<div class="leaderboard-cta">
		<p class="leaderboard-cta__text">{{ text }}</p>
	</div>
</template>

<style lang="scss">
	.leaderboard-cta
	{
		font-size: 12px;
		font-weight: 700;
		padding: 5px 10px;
		border-radius: 14px;
		border: 1px solid $gray;
		background-color: $darkGray;
	}
</style>