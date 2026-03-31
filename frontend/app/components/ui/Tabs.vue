<script setup lang="ts">
	import { type UiTabsHeaderButton } from '~/types/ui/tabs';

	const props = defineProps<{ items: UiTabsHeaderButton[] }>();
	const model = defineModel({ default: 0 });

	const getSwitcherBlockLeftStyle = computed((): string =>
		{
			const tabsLength  = props.items.length || 1;
			const step        = 100 / tabsLength;
			const safeIndex   = Math.max(0, Math.min(model.value, tabsLength - 1));
			const leftPercent = safeIndex * step;

			if (tabsLength === 0)
				return '--width: 0%; --left: 0%; --right: unset;';

			return `
				--width : ${step}%;
				--left  : ${leftPercent}%;
				--right : unset;
			`.trim();
		}
	);

	const changeActiveTab = (idx: number)=> model.value = idx;
</script>

<template>
	<div class="ui-tabs">
		<div
			class="ui-tabs__header"
			:style="getSwitcherBlockLeftStyle"
		>
			<UiButton
				v-for="(button, idx) in items"
				:key="idx"
				color="transparent"
				@click="changeActiveTab(idx)"
				:class="{ 'active': idx === model }"
			>
				{{ button.title }}
			</UiButton>
		</div>

		<div class="ui-tabs__content">
			<slot />
		</div>
	</div>
</template>

<style lang='scss' scoped>
	.ui-tabs
	{
		row-gap: 5px;

		display: flex;
		flex-direction: column;
	}

	.ui-tabs__header
	{
		padding: 2px;
		column-gap: 4px;
		border-radius: 5px;
		background-color: $darkGray;

		display: flex;
		position: relative;

		.ui-button
		{
			z-index: 2;
			flex-basis: var(--width);
			font-weight: 600;
			padding: 10px;

			&.active { color: $primary; }
		}

		&::before
		{
			z-index: 1;
			content: '';
			border-radius: 3px;
			height: calc(100% - 4px);
			width: calc(var(--width) - 4px);

			top: 50%;
			position: absolute;
			right: var(--right);
			background-color: $white;
			left: calc(var(--left) + 2px);
			transform: translateY(-50%);

			@include tr(.3, left, right, width);
		}
	}
</style>