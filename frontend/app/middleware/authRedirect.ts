export default defineNuxtRouteMiddleware((to) => {
	const { token } = useTelegramAuth();

	if (token.value && to.fullPath === '/auth')
		return navigateTo('/');
})