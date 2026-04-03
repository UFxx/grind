<script setup lang="ts">
	import Cup from '~/components/icons/Cup.vue';
	import User from '~/components/icons/User.vue';
	import Home from '~/components/icons/Home.vue';

	const router           = useRoute();
	const { arrivedState } = useScroll(window);

	const onAuthPage = computed(() => router.path === '/auth');

	const activeLink = ref(router.path);

	const handleClick = (link: string) => activeLink.value = link;

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
		},
		{
			icon  : Cup,
			link  : '/leaderboard',
			label : 'Rating',
		}
	];
</script>

<template>
	<Transition name="fade">
		<div
			v-if="!onAuthPage && (!arrivedState.bottom || arrivedState.top)"
			class="footer"
		>
			<NuxtLink
				v-for="link in menuLinks"
				:to="link.link"
				:class="{ active: activeLink === link.link }"
				@click="handleClick(link.link)"
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
		overflow: hidden;
		padding: 5px;
		border-radius: 100px;
		background-color: transparent;

		@include tr(.3s, all);

		&.router-link-active,
		&.active
		{
			color: $primary;
			background-color: $white;

			.footer__item-label
			{
				opacity: 1;
				max-width: 100px;
			}
		}

		&:not(&.router-link-active):not(&.active)
		{
			.footer__item-label
			{
				max-width: 0;
				opacity: 0;
			}
		}
	}

	.footer__item-label
	{
		white-space: nowrap;
		transition: max-width 0.3s ease, opacity 0.3s ease;
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