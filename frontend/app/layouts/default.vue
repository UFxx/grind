<script setup>
	useSeoMeta({ title: 'Grind' })
</script>

<template>
	<div class="default-layout">
		<main class="container">
			<Transition name="fade" mode="out-n">
				<NuxtPage :key="$route.fullPath" />
			</Transition>
		</main>
		<Footer />

		<Toaster />
	</div>

	<PopupWrapper />
</template>

<style lang='scss'>
	.default-layout
	{
		min-height: 100vh;
		background-color: $primary;

		display: flex;
		align-items: center;
		flex-direction: column;

		main { flex-grow: 1; }
	}
</style>