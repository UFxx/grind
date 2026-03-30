<script setup lang="ts">
	import User from '~/components/icons/User.vue';
	import Home from '~/components/icons/Home.vue';

	const router           = useRoute();
	const { arrivedState } = useScroll(window);

	const onAuthPage = computed(() => router.path === '/auth');

	const menuLinks =
	[
		{
			icon  : User,
			link  : '/',
			label : 'Profile',
		},
		{
			icon  : Home,
			link  : '/home',
			label : 'Home',
		}
	];
</script>

<template>
	<Transition name="fade">
		<div
			v-if="!onAuthPage && !arrivedState.bottom"
			class="footer"
		>
			<NuxtLink
				v-for="(link, idx) in menuLinks"
				:key="idx"
				:to="link.link"
				class="footer__item-wr"
			>
				<div class="footer__item">
					<component :is="link.icon" class="footer__item-icon" />
					<span class="footer__item-label">{{ link.label }}</span>
				</div>
			</NuxtLink>
		</div>
	</Transition>
</template>

<style lang='scss' scoped>
	.footer
	{
		padding: 10px;
		column-gap: 10px;
		border-radius: 100px;
		backdrop-filter: blur(4px);
		background-color: rgba(#AAAAAA, 0.1);
		box-shadow: 0 0 6px 2px rgba($gray, 0.25);

		left: 50%;
		bottom: 20px;
		display: flex;
		position: fixed;
		align-items: center;
		transform: translateX(-50%);
	}

	.footer__item-wr
	{
		padding: 5px;
		border-radius: 100px;

		&.router-link-active
		{
			color: $black;
			background-color: $white;
		}

		&:not(&.router-link-active)
		{
			.footer__item-label { display: none; }
		}
	}

	.footer__item
	{
		column-gap: 5px;
		font-size: 12px;
		line-height: 16px;

		display: flex;
		align-items: center;

		@include tr(.3, color);
	}

	.footer__item-icon
	{
		width: 24px;
		height: 24px;
	}
</style>