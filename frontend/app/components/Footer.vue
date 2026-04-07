<script setup lang="ts">
	import Cup from '~/components/icons/Cup.vue';
	import User from '~/components/icons/User.vue';
	import Home from '~/components/icons/Home.vue';

	const router           = useRoute();
	const { arrivedState } = useScroll(window);
	const { togglePopup }  = usePopupsStore();

	const onAuthPage = computed(() => router.path === '/auth');

	const activeLink = ref(router.path);

	const handleClick = (link: string) => activeLink.value = link;

	const openAddActivityPopup = () => togglePopup('AddActivity', true);

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
	<Transition name="scaleY">
		<div
			v-if="!onAuthPage && (!arrivedState.bottom || arrivedState.top)"
			class="footer"
		>
			<div class="footer__menu">
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
			<div
				class="footer__item-wr footer__item-wr--add-activity"
				@click="openAddActivityPopup"
			>
				<div class="footer__item">
					<IconsAdd class="footer__item-icon" />
				</div>
			</div>
		</div>

	</Transition>
</template>

<style lang='scss' scoped>
	.footer
	{
		width: 100%;
		column-gap: 20px;

		left: 50%;
		bottom: 20px;
		display: flex;
		position: fixed;
		align-items: center;
		justify-content: center;
		transform: translateX(-50%);
	}

	.footer__menu
	{
		padding: 10px;
		flex-shrink: 0;
		column-gap: 10px;
		border-radius: 100px;
		backdrop-filter: blur(4px);
		background-color: rgba(#AAAAAA, 0.1);
		box-shadow: 0 0 6px 2px rgba($gray, 0.25);

		display: flex;
		align-items: center;
	}

	.footer__item-wr
	{
		padding: 5px;
		overflow: hidden;
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

		&--add-activity
		{
			padding: 12px;
			flex-shrink: 0;
			cursor: pointer;
			backdrop-filter: blur(4px);
			background-color: rgba(#AAAAAA, 0.1);
			box-shadow: 0 0 6px 2px rgba($gray, 0.25);

			svg
			{
				width: 30px;
				height: 30px;
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
		@include tr(.3, max-width, opacity)
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
		flex-shrink: 0;
	}
</style>