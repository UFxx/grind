<script setup lang="ts">
	import { type UiMiniSwitcher, type UiMiniSwitcherItem } from '~/types/ui/miniSwitcher';

	const props = defineProps<UiMiniSwitcher>();
	const model = defineModel<UiMiniSwitcherItem>({ required: true });

	const isToggled = ref(false);

	const toggleCurrentItem = () =>
	{
		model.value     = props.items.filter(item => item.id !== model.value.id)[0] || model.value;
		isToggled.value = !isToggled.value;
	};
</script>

<template>
	<div class="ui-mini-switcher" @click="toggleCurrentItem">
		<span
			v-for="item in items"
			:key="item.id"
			class="ui-mini-switcher__label"
			:class="{ 'active': item.id === model.id }"
		>
			{{ item.label }}
		</span>

		<div class="ui-mini-switcher__switcher">
			<div
				class="ui-mini-switcher__thumb"
				:class="{ 'ui-mini-switcher__thumb--toggled': isToggled }"
			/>
		</div>
	</div>
</template>

<style lang='scss'>
	.ui-mini-switcher
	{
		column-gap: 5px;
		user-select: none;

		display: flex;
		align-items: center;
	}

	.ui-mini-switcher__label
	{
		color: $gray;
		font-size: 10px;
		cursor: pointer;

		@include tr(.3, color);

		&.active { color: $white; }

		&:last-of-type { order: 1; }
	}

	.ui-mini-switcher__switcher
	{
		width: 25px;
		height: 12px;
		padding: 2px;
		cursor: pointer;
		background-color: $darkGray;
		border-radius: 100px;

		position: relative;
	}

	.ui-mini-switcher__thumb
	{
		width: 8px;
		height: 8px;
		border-radius: 100%;
		background-color: $white;

		@include tr(.3, left, transform);

		left: 2px;
		position: absolute;
		transform: translateX(0);

		&--toggled
		{
			left: calc(100% - 2px);
			transform: translateX(-100%);
		}
	}
</style>