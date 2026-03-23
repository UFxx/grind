<script setup lang="ts">
	const props = withDefaults(defineProps<{
			value     : number | string;
			delay?    : number;
			duration? : number;
			prefix?   : string;
			suffix?   : string;
		}>(),
		{
			delay    : 0,
			duration : 1000,
			prefix   : '',
			suffix   : ''
		}
	);

	const displayValue = ref(0);
	const isAnimated   = ref(false);

	const numericValue = computed(() =>
	{
		const num = Number(props.value);
		return isNaN(num) ? 0 : num;
	});

	const formattedValue = computed(() =>
	{
		const isFloat = String(props.value).includes('.');

		return isFloat
			? displayValue.value.toFixed(1)
			: Math.round(displayValue.value);
	});

	onMounted(() =>
		{
			setTimeout(() => animateCount(), props.delay);
		}
	);

	const animateCount = () =>
	{
		const start     = 0;
		const end       = numericValue.value;
		const startTime = performance.now();

		const animate = (currentTime: number) => {
			const elapsed = currentTime - startTime;
			const progress = Math.min(elapsed / props.duration, 1);

			const ease = 1 - (1 - progress) * (1 - progress);

			displayValue.value = start + (end - start) * ease;

			if (progress < 1)
				requestAnimationFrame(animate);
			else
			{
				displayValue.value = end;
				isAnimated.value = true;
			}
		};

		requestAnimationFrame(animate);
	};
</script>

<template>
	<span class="ui-count-number">
		{{ prefix }}{{ formattedValue }}{{ suffix }}
	</span>
</template>

<style lang='scss' scoped>
	.ui-count-number { display: inline-block; }
</style>