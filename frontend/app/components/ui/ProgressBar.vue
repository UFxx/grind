<script setup lang="ts">
	import { type UiProgressBar } from '~/types/ui/progressBar';

	withDefaults(defineProps<UiProgressBar>(), { color: 'green', });
</script>

<template>
	<div
		class="ui-progress-bar"
		:class="`ui-progress-bar--${color}`"
		:style="`--progress: ${progress}%`"
	/>
</template>

<style lang='scss'>
	.ui-progress-bar
	{
		height: 2px;
		width: 100%;
		border-radius: 100px;
		background-color: $darkGray;

		position: relative;

		&::after
		{
			content: '';
			border-radius: 100px;
			width: var(--progress);

			@include tr(0.3, width);

			position: absolute;
			height: 100%;
		}

		&--green::after
		{
			background-color: $green;
			box-shadow: 0 0 6px 1px rgba($green, 0.5);
		}

		&--blue::after
		{
			background-color: $blue;
			box-shadow: 0 0 6px 1px rgba($blue, 0.5);
		}
	}
</style>