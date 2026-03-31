<script setup lang="ts">
	import { type UiTextarea } from '~/types/ui/textarea';

	const props = withDefaults(defineProps<UiTextarea>(),
		{
			fullWidth   : false,
			placeholder : '',
		}
	);

	const model = defineModel<string>();
</script>

<template>
	<textarea
		class="ui-textarea"
		:class="{ 'ui-textarea--full-width': fullWidth }"
		:placeholder
		v-model="model"
	/>
</template>

<style lang='scss' scoped>
	.ui-textarea
	{
		resize: none;
		font-size: 12px;
		font-weight: 500;
		min-height: 100px;
		padding: 7px 10px;
		border-radius: 5px;
		background-color: #0C0C0C;
		outline: 1px solid $darkGray;

		&--full-width { width: 100%; }
	}
</style>