<script setup>
	import WebApp from '@twa-dev/sdk';

	const userStore       = useUserStore();
	const { togglePopup } = usePopupsStore();
	const { user }        = useApi();

	const inviter = userStore.userData.inviter;

	const closeUserInfoPopup        = () => togglePopup('UserInfo', false);
	const openFeedbackPopup = () =>
	{
		closeUserInfoPopup();
		togglePopup('Feedback', true);
	};
	const deleteAccount     = async () =>
	{
		try
		{
			const response = await user.deleteAccount()

			if (!response.data.length)
				WebApp.close();
		}
		catch(err) { console.error(err) }
	};
</script>

<template>
	<div class="user-info-popup-wr">
		<IconsClose class="user-info-popup__close" @click="closeUserInfoPopup" />
		<p class="user-info-popup__title">User information</p>

		<div class="user-info-popup__content">
			<div class="user-info-popup__content-inviter">
				<p class="user-info-popup__content-inviter-title">Inviter</p>
				<div class="user-info-popup__content-inviter-content">
					<img
						class="user-info-popup__content-inviter-avatar"
						:src="inviter.avatarUrl"
						:alt="`${inviter.name} avatar`"
					/>
					<div class="user-info-popup__content-inviter-info">
						<p class="user-info-popup__content-inviter-tag">@{{ inviter.name }}</p>
						<p class="user-info-popup__content-inviter-lvl-wr">
							<span>{{ inviter.level.currentLevel }}</span>
							<span class="user-info-popup__content-inviter-lvl-label">lvl</span>
						</p>
					</div>
				</div>
			</div>
		</div>
		<div class="user-info-popup__footer">
			<UiButton
				color="red"
				:fullWidth="true"
				@click="deleteAccount"
			>
				Delete account
			</UiButton>
			<UiButton
				color="white"
				:fullWidth="true"
				@click="openFeedbackPopup"
			>
				Leave a feedback
			</UiButton>
		</div>
	</div>
</template>

<style lang='scss' scoped>
	.user-info-popup-wr
	{
		width: 100%;
		row-gap: 10px;
		padding: 10px;
		border-radius: 10px;
		background-color: $primary;
		box-shadow: 0 0 6px 2px $darkGray;

		display: flex;
		position: relative;
		flex-direction: column;
	}

	.user-info-popup__close
	{
		opacity: 0.5;
		cursor: pointer;
		user-select: none;

		top: 10px;
		right: 10px;
		position: absolute;

		@include tr(.3, opacity);

		&:hover { opacity: 1; }
	}

	.user-info-popup__title
	{
		font-weight: 700;
		text-align: center;
		padding-bottom: 10px;
		border-bottom: 2px solid $darkGray;
	}

	.user-info-popup__content-inviter
	{
		row-gap: 10px;

		display: flex;
		flex-direction: column;
	}

	.user-info-popup__content-inviter-title
	{
		font-size: 12px;
		font-weight: 500;
	}

	.user-info-popup__content
	{
		padding-bottom: 10px;
		border-bottom: 2px solid $darkGray;
	}

	.user-info-popup__content-inviter-avatar
	{
		width: 30px;
		height: 30px;
		border-radius: 100%;
		border: 1px solid $white;
		border-spacing: -1px;
	}

	.user-info-popup__content-inviter-content
	{
		column-gap: 10px;

		display: flex;
		align-items: center;
	}

	.user-info-popup__content-inviter-tag,
	.user-info-popup__content-inviter-lvl-wr
	{
		font-size: 12px;
		line-height: 16px;
	}

	.user-info-popup__content-inviter-lvl-label { color: $gray; }

	.user-info-popup__footer
	{
		display: flex;
		column-gap: 10px;
	}
</style>