<script setup lang="ts">
	const toastsStore = useToastStore();

	const closeToast = (id: number) => toastsStore.removeToast(id);
</script>

<template>
	<TransitionGroup
		class="toaster"
		name="toaster"
		tag="div"
	>
		<UiToast
			v-for="toast in toastsStore.toasts"
			:key="toast.id"
			:toast="toast"
			@close="closeToast"
		/>
	</TransitionGroup>
</template>

<style lang="scss">
.toaster
{
	row-gap: 10px;

	top: 10px;
	right: 10px;
	display: flex;
	position: fixed;
	flex-direction: column;
	align-items: flex-end;
}
</style>