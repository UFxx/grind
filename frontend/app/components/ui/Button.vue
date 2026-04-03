<script setup lang="ts">
	import { type UiButton } from '~/types/ui/button';

	withDefaults(defineProps<UiButton>(),
		{
			color: 'black',
			disabled: false,
			fullWidth: false
		}
	);
</script>

<template>
	<button
		class="ui-button"
		:class="[
			`ui-button--${color}`,
			{ 'ui-button--full-width': fullWidth }
		]"
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
		font-size: 12px;
		min-height: 26px;
		line-height: 16px;
		padding: 5px 10px;
		text-align: center;
		border-radius: 3px;
		background-color: $primary;

		@include tr(.3, opacity, background-color, color);

		&:hover { opacity: 0.9; }
		&:disabled
		{
			background-color: $gray;
			color: $white;
		}

		&--white
		{
			color: $primary;
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
			background-color: $primary;
		}

		&--red
		{
			color: $white;
			background-color: rgba($red, 0.3);
		}

		&--full-width { width: 100%; }
	}
</style>