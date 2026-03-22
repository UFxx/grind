<script setup lang="ts">
	import { type Toast } from '~/types/toast'

	const props = defineProps<{ toast: Toast }>();

	const emit = defineEmits(['close']);

	const closeToast = () => emit("close", props.toast.id);
</script>

<template>
	<div
		class="ui-toast"
		:class="`ui-toast--${toast.type}`"
		@click="closeToast"
	>
		<p class="ui-toast__text">{{ toast.text }}</p>
	</div>
</template>

<style lang="scss">
	.ui-toast
	{
		max-width: 240px;
		padding: 5px 10px;
		border-radius: 4px;

		&--error { background-color: $red; }
		&--success { background-color: $green; }
		&--info { background-color: $orange; }

		@include mq($tablet) { max-width: 280px; }
		@include mq($desktop) { max-width: 320px; }
		@include mq($wide) { max-width: 360px; }
	}

	.ui-toast__text
	{
		color: $white;
		user-select: none;
		text-align: center;
		overflow-wrap: break-word;
	}
</style>