<script setup lang="ts">
	import { type Toast } from '~/types/toast'

	const props = defineProps<{toast: Toast}>();

	const emit = defineEmits(['close']);

	const closeToast = () => emit("close", props.toast.id)
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
	border-radius: 4px;
	padding: 5px 10px;
	max-width: 240px;

	&--error { background-color: $lightRed; }
	&--success { background-color: $lightGreen; }
	&--info { background-color: $lightOrange; }

	@media screen and (min-width: $tabletScreen) { max-width: 280px; }
	@media screen and (min-width: $desktopScreen) { max-width: 320px; }
	@media screen and (min-width: $wideScreen) { max-width: 360px; }
}

.ui-toast__text
{
	color: white;
	text-align: center;
	user-select: none;
	overflow-wrap: break-word;
}
</style>