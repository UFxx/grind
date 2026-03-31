<template>
	<div class="ui-loader"></div>
</template>

<style lang='scss'>
	.ui-loader
	{
		width: 24px;
		height: 24px;
		margin: auto;
		position: relative;
	}

	.ui-loader:after
	{
		content: '';
		width: 100%;
		height: 100%;
		background: $primary;
		animation: bxSpin 0.5s linear infinite;
		position: absolute;
		top: -8px;
		left: 0;
		border-radius: 4px;
	}

	@keyframes bxSpin
	{
		17% { border-bottom-right-radius: 3px; }
		25% { transform: translateY(5px) rotate(22.5deg); }
		50%
{
			transform: translateY(10px) scale(1, .9) rotate(45deg);
			border-bottom-right-radius: 40px;
		}
		75% { transform: translateY(5px) rotate(67.5deg); }
		100% { transform: translateY(0) rotate(90deg); }
	}

	@keyframes shadow
	{
		0%, 100% { transform: scale(1, 1); }
		50% { transform: scale(1.2, 1); }
	}
</style>