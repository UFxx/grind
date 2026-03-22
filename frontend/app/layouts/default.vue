<script setup></script>

<template>
	<div class="default-layout">
		<main>
			<slot />
		</main>
		<Toaster />
	</div>
</template>

<style lang='scss'>
.default-layout
{
	display: flex;
	min-height: 100vh;
	background: black;
	flex-direction: column;

	main { flex-grow: 1; }
}
</style>