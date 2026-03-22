<script setup lang="ts">
	import { type UiButton } from '~/types/ui/button';

	withDefaults(defineProps<UiButton>(),
		{
			color: 'black',
			disabled: false
		}
	);
</script>

<template>
	<button
		class="ui-button"
		:class="`ui-button--${color}`"
		:disabled
	>
		<slot />
	</button>
</template>

<style lang='scss' scoped>
	.ui-button
	{
		opacity: 1;
		color: $white;
		padding: 10px;
		font-size: 12px;
		line-height: 16px;
		text-align: center;
		border-radius: 3px;
		background-color: $black;
		@include tr(.3s, opacity, background-color, color);

		&:hover { opacity: 0.9; }
		&:disabled
		{
			background-color: $gray;
			color: $white;
		}

		&--white
		{
			color: $black;
			background-color: $white;
		}

		&--transparent
		{
			color: $white;
			background-color: transparent;

			&:hover { opacity: unset; }
		}

		&--black
		{
			color: $white;
			background-color: $black;
		}
	}
</style>