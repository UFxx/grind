<script setup lang="ts">
	const toastStore = useToastStore();

	const closeToast = (id: number) => toastStore.removeToast(id)
</script>

<template>
	<TransitionGroup
		class="toaster"
		name="toaster"
		tag="div"
	>
		<UiToast
			v-for="toast in toastStore.getToasts()"
			:toast="toast"
			:key="toast.id"
			@close="closeToast"
		/>
	</TransitionGroup>
</template>

<style lang="scss">
.toaster
{
	display: flex;
	flex-direction: column;
	align-items: flex-end;

	position: fixed;
	top: 10px;
	right: 10px;

	row-gap: 10px;
}

.toaster-move,
.toaster-enter-active,
.toaster-leave-active {
  transition: all 0.5s ease;
}

.toaster-enter-from,
.toaster-leave-to {
  opacity: 0;
  transform: translateX(100vw);
}

.toaster-leave-active {
  position: absolute;
}
</style>