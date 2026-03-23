<script setup lang="ts">
	import { type UiProgressBar } from '~/types/ui/progressBar';

	const props = withDefaults(defineProps<UiProgressBar>(), { color: 'green', });

	const isVisible = ref(false);

	setTimeout(() => isVisible.value = true, props.delay ?? 300);
</script>

<template>
	<div
		class="ui-progress-bar"
		:class="`ui-progress-bar--${color}`"
	>
		<div
			class="ui-progress-bar__fill"
			:class="{ 'ui-progress-bar__fill--visible': isVisible }"
			:style="`--target-width: ${progress}%`"
		/>
	</div>
</template>

<style lang='scss'>
	.ui-progress-bar
	{
		height: 2px;
		width: 100%;
		border-radius: 100px;
		background-color: $darkGray;

		position: relative;

		&__fill
		{
			width: 0%;
			height: 100%;
			border-radius: 100px;
			transition: width 0.6s cubic-bezier(0.4, 0, 0.2, 1);

			top: 0;
			left: 0;
			position: absolute;

			&--visible { width: var(--target-width); }
		}

		&--green &__fill
		{
			background-color: $green;
			box-shadow: 0 0 6px 1px rgba($green, 0.5);
		}

		&--blue &__fill
		{
			background-color: $blue;
			box-shadow: 0 0 6px 1px rgba($blue, 0.5);
		}
	}
</style>