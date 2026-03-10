<script setup></script>

<template>
	<div class="default-layout">
		<Header />
		<main>
			<slot />
		</main>
		<Footer />
	</div>
</template>

<style lang='scss'>
.default-layout
{
	display: flex;
	flex-direction: column;
	min-height: 100vh;

	main { flex-grow: 1; }
}
</style>