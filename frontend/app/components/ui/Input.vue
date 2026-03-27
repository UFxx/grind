<script setup lang="ts">
	import { type UiInput } from '~/types/ui/input';

	const props = withDefaults(defineProps<UiInput>(),
		{
			type        : 'text',
			inputmode   : 'text',
			fullWidth   : false,
			placeholder : '',
			onlyNumbers : false
		}
	);

	const model = defineModel<string>();

	const keydownHandler = (e: KeyboardEvent) =>
	{
		if (!props.onlyNumbers) return;

		const allowedKeys = ['Backspace', 'Delete', 'ArrowLeft', 'ArrowRight', 'Tab', 'Home', 'End'];
		if (allowedKeys.includes(e.key))
			return;

		if (e.key >= '0' && e.key <= '9')
			return;

		e.preventDefault();
	};

	const inputHandler = (e: InputEvent) =>
	{
		const input = e.target as HTMLInputElement;

		if (props.onlyNumbers)
		{
			const sanitized = input.value.replace(/\D/g, '');
			model.value = sanitized;
		}
		else model.value = input.value;
	};
</script>

<template>
	<input
		class="ui-input"
		:class="{ 'ui-input--full-width': fullWidth }"
		:type
		:inputmode
		:placeholder
		@keydown="keydownHandler"
		@input="inputHandler"
		:value="model"
	>
</template>

<style lang='scss'>
	.ui-input
	{
		font-size: 12px;
		font-weight: 500;
		padding: 7px 10px;
		border-radius: 5px;
		background-color: #0C0C0C;
		outline: 1px solid $darkGray;

		&--full-width { width: 100%; }
	}
</style>