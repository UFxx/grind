<script setup>
	const nuxtApp = useNuxtApp();

	useSeoMeta({ title: 'Grind' });

	nuxtApp.hook("page:finish", () => getScrollWidth());
</script>

<template>
	<div class="default-layout">
		<main class="container">
			<Transition name="fade" mode="out-in">
				<NuxtPage />
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