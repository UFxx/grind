<script setup>
	useSeoMeta({ title: 'Grind' })
</script>

<template>
	<div class="default-layout">
		<main class="container">
			<slot />
		</main>
	</div>
</template>

<style lang='scss'>
	.default-layout
	{
		min-height: 100vh;
		background: $black;

		display: flex;
		align-items: center;
		flex-direction: column;

		main
		{
			flex-grow: 1;
			font-family: $Manrope;
		}
	}
</style>