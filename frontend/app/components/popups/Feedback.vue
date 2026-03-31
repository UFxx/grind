<script setup lang="ts">
	import { type Feedback } from '~/types/feedback';

	const { togglePopup } = usePopupsStore();
	const { other }       = useApi();
	const { addToast }    = useToastsStore();
	const userStore       = useUserStore();

	const choosedRating = ref(0);
	const review        = ref('');
	const isLoading     = ref(false);

	const closeFeedbackPopup = () => togglePopup('Feedback', false);

	const setRating = (value: number) => {
		if (choosedRating.value === value)
		{
			choosedRating.value = 0;
			return;
		}

		choosedRating.value = value
	};

	const sendFeedback = async () =>
	{
		const payload: Feedback =
		{
			sender  : userStore.userData?.name || '',
			rating  : choosedRating.value,
			message : review.value
		};

		isLoading.value = true;

		try
		{
			const response = await other.sendFeedback(payload);

			if (!response.data.length)
			{
				closeFeedbackPopup();
				addToast('success', 'Thanks for your feedback!');
			}
		}
		catch (err)
		{
			console.error(err);
			addToast('error', 'There was an error sending your review');
		}
		finally { isLoading.value = false; }
	};
</script>

<template>
	<div class="feedback-popup-wr">
		<IconsClose class="feedback-popup__close" @click="closeFeedbackPopup" />
		<p class="feedback-popup__title">Leave a feedback</p>

		<div class="feedback-popup__content">
			<div class="feedback-popup__rate feedback-popup__section">
				<p class="feedback-popup__section-title">Please rate our app</p>
				<div class="feedback-popup__rate-content">
					<IconsStar
						v-for="(_, idx) in 5"
						:key="idx"
						class="feedback-popup__rate-content-item"
						:class="{ 'active': idx < choosedRating }"
						@click="setRating(idx + 1)"

						v-motion-fade
						:duration="300"
						:delay="idx * 100"
					/>
				</div>
			</div>

			<div class="feedback-popup__review feedback-popup__section">
				<p class="feedback-popup__section-title">...and leave a review</p>
				<div class="feedback-popup__review-content">
					<UiTextarea
						:fullWidth="true"
						placeholder="Enter your review"
						v-model="review"
					/>
				</div>
			</div>
		</div>

		<UiButton
			color="white"
			:fullWidth="true"
			:disabled="isLoading"
			@click="sendFeedback"
		>
			Send
		</UiButton>
	</div>
</template>

<style lang='scss'>
	.feedback-popup-wr
	{
		width: 100%;
		row-gap: 20px;
		padding: 10px;
		border-radius: 10px;
		background-color: $primary;
		box-shadow: 0 0 6px 2px $darkGray;

		display: flex;
		position: relative;
		flex-direction: column;
	}

	.feedback-popup__close
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

	.feedback-popup__title
	{
		font-weight: 700;
		text-align: center;
	}

	.feedback-popup__content
	{
		row-gap: 20px;

		display: flex;
		flex-direction: column;
	}

	.feedback-popup__section
	{
		row-gap: 10px;

		display: flex;
		flex-direction: column;
	}

	.feedback-popup__section-title
	{
		font-size: 12px;
		line-height: 16px;
	}

	.feedback-popup__rate-content
	{
		padding: 0 10px;
		overflow: hidden;

		display: flex;
		column-gap: 10px;
	}

	.feedback-popup__rate-content-item
	{
		color: $gray;
		cursor: pointer;

		@include tr(.3, color);

		&.active { color: $white; }
	}
</style>