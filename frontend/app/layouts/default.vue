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
		display: flex;
		align-items: center;
		flex-direction: column;

		min-height: 100vh;
		background: $black;

		main
		{
			flex-grow: 1;
			font-family: $Manrope;
		}
	}
</style>