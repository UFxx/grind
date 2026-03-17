<script setup lang="ts">
	defineProps(
		{
			color:
			{
				type: String,
				required: false,
				default: 'black',
				validator: (v :string) => ['white', 'black', 'transparent'].includes(v)
			}
		}
	)
</script>

<template>
	<button
		class="ui-button"
		:class="`ui-button--${color}`"
	>
		<slot />
	</button>
</template>

<style lang='scss' scoped>
	.ui-button
	{
		color: $white;
		padding: 10px;
		font-size: 12px;
		line-height: 16px;
		text-align: center;
		border-radius: 3px;
		background-color: $black;
		@include tr(.3s, opacity);

		&:hover { opacity: 0.9; }

		&--white
		{
			color: $black;
			background-color: $white;
		}

		&--transparent
		{
			color: $white;
			background-color: transparent;

			&:hover { opacity: unset; }
		}

		&--black
		{
			color: $white;
			background-color: $black;
		}
	}
</style>