<script setup lang="ts">
	import { usePopupsRegistry } from '~/composables/usePopupsRegistry';

	const popupsStore = usePopupsStore();

	const activeComponent = computed(() => usePopupsRegistry(popupsStore.activePopup));
</script>

<template>
	<Teleport to="#teleports">
		<Transition name="opacity">
			<div v-if="activeComponent" class="popup-wr container">
				<component :is="activeComponent" />
			</div>
		</Transition>
	</Teleport>
</template>

<style lang='scss'>
	.popup-wr
	{
		z-index: 2;
		width: 100vw;
		height: 100dvh;
		padding-top: 0;
		padding-bottom: 0;
		backdrop-filter: blur(8px);
		background-color: rgba($black, $alpha: 0.9);

		top: 50%;
		left: 50%;
		display: flex;
		position: fixed;
		align-items: center;
		justify-content: center;
		transform: translate(-50%, -50%);
	}
</style>